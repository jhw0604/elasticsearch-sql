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

type elasticsearchQueryRequest struct {
	Query string `json:"query"`
}

type elasticsearchCursorRequest struct {
	Cursor string `json:"cursor"`
}

type elasticsearchColumnInfo struct {
	Name string            `json:"name"`
	Type elasticsearchType `json:"type"`
}

type elasticsearchResult struct {
	Columns []elasticsearchColumnInfo `json:"columns"`
	Rows    [][]interface{}           `json:"rows"`
	Cursor  string                    `json:"Cursor"`
}

func newRows(dsn string, query string) *Rows {
	byteReqBody, err := json.Marshal(elasticsearchQueryRequest{query})
	if err != nil {
		panic(err)
	}

	return elasticsearchRequest(dsn, string(byteReqBody))
}

func nextRows(dsn string, cursor string) ([][]driver.Value, string) {
	byteReqBody, err := json.Marshal(elasticsearchCursorRequest{cursor})
	if err != nil {
		panic(err)
	}
	result := elasticsearchRequest(dsn, string(byteReqBody))

	return (*result).rows, (*result).cursor
}

func parsingDSN(dsn string) (url, username, password string) {
	var protocal, address, port string

	dnsParts := strings.Split(dsn, "://")
	if len(dnsParts) <= 1 {
		protocal = "http"
		dsn = dnsParts[0]
	} else {
		protocal = dnsParts[0]
		dsn = dnsParts[1]
	}

	var certBase64 string
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
			panic(err)
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

	return protocal + "://" + address + ":" + port + "/_xpack/sql", username, password
}

func postElasticsearch(dsn string, body string) string {
	url, username, password := parsingDSN(dsn)

	client := http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-type", "application/json")
	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var elaResp string
	buff := make([]byte, 10)
	n, err := io.ReadFull(resp.Body, buff)
	for err != io.EOF {
		elaResp = elaResp + string(buff[:n])
		n, err = io.ReadFull(resp.Body, buff)
	}

	return elaResp
}

func elasticsearchRequest(dsn string, body string) *Rows {
	elaResp := postElasticsearch(dsn, body)
	elaResult := elasticsearchResult{}

	err := json.Unmarshal([]byte(elaResp), &elaResult)
	if err != nil {
		panic(err)
	}

	if elaResult.Rows == nil {
		panic(errors.New(elaResp))
	}

	var columns []string
	var types []elasticsearchType
	for _, columnInfo := range elaResult.Columns {
		columns = append(columns, columnInfo.Name)
		types = append(types, columnInfo.Type)
	}

	var rows [][]driver.Value
	for _, values := range elaResult.Rows {
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
		cursor:  elaResult.Cursor,
	}
}

func typeConvert(elaType elasticsearchType, value interface{}) driver.Value {
	//Unsupported
	//ElaTypeBinary, ElaTypeByte, ElaTypeObject, ElaTypeNested, ElaTypeUnsupported
	switch elaType {
	case ElaTypeKeyword, ElaTypeText, ElaTypeIP:
		return value.(string)
	case ElaTypeShort, ElaTypeLong, ElaTypeFloat, ElaTypeHalfFloat, ElaTypeScaledFloat, ElaTypeDouble:
		return value.(float64)
	case ElaTypeInteger:
		return int(value.(float64))
	case ElaTypeBoolean:
		return value.(bool)
	case ElaTypeDatetime:
		t, err := time.Parse(time.RFC3339, value.(string))
		if err != nil {
			return nil
		}
		return t
	case ElaTypeNull:
		return nil
	default:
		return value
	}

}
