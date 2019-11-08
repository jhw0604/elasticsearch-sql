package elasticsearch

import (
	"bytes"
	"database/sql/driver"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type esQueryRequest struct {
	Query string `json:"query"`
}

type esCursorRequest struct {
	Cursor string `json:"cursor"`
}

type esColumnInfo struct {
	Name string `json:"name"`
	Type esType `json:"type"`
}

type esResult struct {
	Columns []esColumnInfo  `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
	Cursor  string          `json:"Cursor"`
}

func newRows(dsn string, query string) (*Rows, error) {
	byteReqBody, err := json.Marshal(esQueryRequest{query})
	if err != nil {
		return nil, err
	}

	return esRequest(dsn, string(byteReqBody))
}

func nextRows(dsn string, cursor string) ([][]driver.Value, string, error) {
	byteReqBody, err := json.Marshal(esCursorRequest{cursor})
	if err != nil {
		return nil, "", err
	}
	result, err := esRequest(dsn, string(byteReqBody))

	return (*result).rows, (*result).cursor, err
}

func parsingDSN(dsn string) (url, username, password string, err error) {
	var protocal, address, port, certBase64 string

	dnsParts := strings.Split(dsn, "://")
	if len(dnsParts) <= 1 {
		protocal = "http"
		dsn = dnsParts[0]
	} else {
		protocal = dnsParts[0]
		dsn = dnsParts[1]
	}

	dnsParts = strings.Split(dsn, "@")
	if len(dnsParts) <= 1 {
		certBase64 = ""
		dsn = dnsParts[0]
	} else {
		certBase64 = dnsParts[0]
		dsn = dnsParts[1]
	}

	if certBase64 != "" {
		certByte, err := base64.URLEncoding.DecodeString(certBase64)
		if err != nil {
			return "", "", "", ErrInvalidArgs
		}
		certPart := strings.Split(string(certByte), ":")
		username, password = certPart[0], certPart[1]
	}

	dnsParts = strings.Split(dsn, ":")
	if len(dnsParts) <= 1 {
		address = "localhost"
		port = "9200"
	} else {
		address = dnsParts[0]
		if len(dnsParts[1]) == 0 {
			port = "9200"
		} else {
			port = dnsParts[1]
		}
	}

	return protocal + "://" + address + ":" + port + "/_xpack/sql", username, password, nil
}

func postEs(dsn string, body string) (string, error) {
	url, username, password, err := parsingDSN(dsn)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-type", "application/json")
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var esResp string
	buff := make([]byte, 10)
	n, err := io.ReadFull(resp.Body, buff)
	for err != io.EOF {
		esResp = esResp + string(buff[:n])
		n, err = io.ReadFull(resp.Body, buff)
	}

	return esResp, nil
}

func esRequest(dsn string, body string) (*Rows, error) {
	esResp, err := postEs(dsn, body)
	if err != nil {
		return nil, err
	}

	esResult := esResult{}
	err = json.Unmarshal([]byte(esResp), &esResult)
	if err != nil {
		return nil, err
	}

	if esResult.Rows == nil {
		return nil, errors.New(esResp)
	}

	var columns []string
	var types []esType
	for _, columnInfo := range esResult.Columns {
		columns = append(columns, columnInfo.Name)
		types = append(types, columnInfo.Type)
	}

	var rows [][]driver.Value
	for _, values := range esResult.Rows {
		var row []driver.Value
		for i, value := range values {
			row = append(row, typeConvert(types[i], value))
		}
		rows = append(rows, row)
	}

	return &Rows{
			dsn:     dsn,
			columns: columns,
			types:   types,
			rows:    rows,
			cursor:  esResult.Cursor,
		},
		nil
}

func typeConvert(t esType, value interface{}) driver.Value {
	//Unsupported
	//esBinary, esByte, esObject, esNested, esUnsupported
	switch t {
	case esKeyword, esText, esIP:
		return value.(string)
	case esShort, esLong, esFloat, esHalfFloat, esScaledFloat, esDouble:
		return value.(float64)
	case esInteger:
		return int(value.(float64))
	case esBoolean:
		return value.(bool)
	case esDatetime:
		t, err := time.Parse(time.RFC3339, value.(string))
		if err != nil {
			return nil
		}
		return t
	case esNull:
		return nil
	default:
		return value
	}

}
