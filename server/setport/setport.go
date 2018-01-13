package setport

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

//ポート情報
type (
	portInfo struct {
		Serial_no    string `db:"serial_no" json:"serial_no"`	//シリアルNo(引数)
		Port_no    	 int `db:"port_no" json:"port_no"`				//ポート番号(引数)
		Value    		 string `db:"value" json:"value"`					//設定値(引数)
	}
)

var (
	tablename = "t_port"	//接続ポートテーブル
	conn, _   = dbr.Open("mysql", "iot:@tcp(127.0.0.1:3306)/smw", nil)
	sess      = conn.NewSession(nil)
)

//ポート情報更新
func SetPort(c echo.Context) error {
	port := new(portInfo)

	//ポートエラーを拾う
	if err := c.Bind(port); err != nil {
		return err
	}

	//接続ポートテーブルのupdateSQLを発行する
  conn.Exec(`
		UPDATE t_port
		SET value = "on"
		WHERE serial_no = ? , port_no = ?` ,
		port.Serial_no ,
		port.Port_no ,
	)
	return c.JSON(http.StatusOK , "あっぷでーと")
}