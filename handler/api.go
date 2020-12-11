package handler

import (
	"echo/server"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type menu struct {
	Id_menu    string
	Nama_menu  string
	Deksripsi  string
	Jenis      string
	Harga      string
	Url_gambar string
	Total_order string
}

var data []menu

func BacaData(c echo.Context) error {
	menu_makanan()
	return c.JSON(http.StatusOK, data)

}
func BacaPopuler(c echo.Context) error {
	menu_populer()
	return c.JSON(http.StatusOK, data)

}

func TambahData(c echo.Context) error {
	db, err := server.Koneksi()
	defer db.Close()
	var nama = c.FormValue("Nama_menu")
	var deksripsi = c.FormValue("Deksripsi")
	var harga = c.FormValue("Harga")
	var jenis = c.FormValue("Jenis")
	var url_gambar = c.FormValue("Url_gambar")

	_, err = db.Exec("insert into tbl_menu values (?,?,?,?,?,?)", nil,
		nama, deksripsi, url_gambar, jenis, harga)
	if err != nil {
		fmt.Println("menu gagal ditambahkan")
		return c.JSON(http.StatusOK, "Gagal menambahkan menu")

	} else {
		fmt.Println("menu berhasil ditambahkan")
		return c.JSON(http.StatusOK, "Berhasil menambahkan menu")
	}
}
func InputOrder(c echo.Context) error {
	db, err := server.Koneksi()
	defer db.Close()
	var id = c.FormValue("id")
	var nama_pemesan = c.FormValue("nama_pemesan")
	var nomor_telepon = c.FormValue("nomor_telepon")
	var jumlah = c.FormValue("jumlah")
	var alamat = c.FormValue("alamat")
	
	_, err = db.Exec("insert into tbl_order values (?,?,?,?,?,?)", nil,
	id, nama_pemesan, nomor_telepon, alamat , jumlah)
	if err != nil {
		fmt.Println("Pesanan gagal ditambahkan")
		return c.HTML(http.StatusOK, "<strong>Gagal menambahkan Pesanan</strong>")

	} else {
		fmt.Println("Order berhasil ditambahkan")
		return c.HTML(http.StatusOK, "<script>alert('Berhasil melakukan pemesanan silahkan tunggu telepon dari kami.... Terima Kasi!'); window.location='http://localhost:1323';</script>")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func UbahData(c echo.Context) error {
	db, err := server.Koneksi()
	defer db.Close()

	var id = c.FormValue("Id_menu")
	var nama = c.FormValue("Nama_menu")
	var deksripsi = c.FormValue("Deksripsi")
	var harga = c.FormValue("Harga")
	var jenis = c.FormValue("Jenis")
	var url_gambar = c.FormValue("Url_gambar")

	_, err = db.Exec("update tbl_menu set nama_menu = ? , deksripsi = ? , harga = ? , jenis = ? , url_gambar = ? where id_menu = ?",
	 nama, deksripsi, harga, jenis, url_gambar, id)
	if err != nil {
		fmt.Println("menu gagal diubah")
		return c.JSON(http.StatusOK, "Gagal mengubah menu")

	} else {
		fmt.Println("menu berhasil diubah")
		return c.JSON(http.StatusOK, "Berhasil mengubah menu")
	}
}
func HapusData(c echo.Context) error {
	db, err := server.Koneksi()
	defer db.Close()

	var id = c.FormValue("Id_menu")


	_, err = db.Exec("delete from tbl_menu where id_menu = ?", id)
	if err != nil {
		fmt.Println("menu gagal dihapus")
		return c.JSON(http.StatusOK, "Gagal menghapus menu")

	} else {
		fmt.Println("menu berhasil dihapus")
		return c.JSON(http.StatusOK, "Berhasil mengapus menu")
	}
}

func menu_makanan() {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from tbl_menu")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu,
			&each.Deksripsi, &each.Url_gambar, &each.Jenis,
			&each.Harga)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
		fmt.Println(data)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

}

func menu_populer() {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from vw_totalorder order by total_order Desc limit 8")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu,
			&each.Deksripsi, &each.Url_gambar, &each.Jenis,
			&each.Harga, &each.Total_order)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
		fmt.Println(data)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

}