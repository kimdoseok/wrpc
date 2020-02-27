// iits.go
package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Iit struct {
	Iit_code            string
	Iit_ict_code        sql.NullString
	Iit_pvd_code        sql.NullString
	Iit_ium_code        sql.NullString
	Iit_ity_code        sql.NullString
	Iit_ium_part        sql.NullString
	Iit_iwh_code        sql.NullString
	Iit_ytx_code        sql.NullString
	Iit_ocl_code        sql.NullString
	Iit_ycp_code        sql.NullString
	Iit_opr_code        sql.NullString
	Iit_oct_code        sql.NullString
	Iit_ist_code        sql.NullString
	Iit_ean13           sql.NullString
	Iit_ean13a          sql.NullString
	Iit_name            sql.NullString
	Iit_namealt         sql.NullString
	Iit_vendor_sku      sql.NullString
	Iit_package         sql.NullString
	Iit_qty_base        sql.NullInt64
	Iit_qty_oh          sql.NullInt64
	Iit_pqty_oh         sql.NullInt64
	Iit_qty_bk          sql.NullInt64
	Iit_pqty_bk         sql.NullInt64
	Iit_qty_po          sql.NullInt64
	Iit_pqty_po         sql.NullInt64
	Iit_qty_it          sql.NullInt64
	Iit_pqty_it         sql.NullInt64
	Iit_qty_al          sql.NullInt64
	Iit_pqty_al         sql.NullInt64
	Iit_initqty         sql.NullInt64
	Iit_qdiscqty        sql.NullInt64
	Iit_point           sql.NullInt64
	Iit_leaddays        sql.NullInt64
	Iit_minlevel        sql.NullInt64
	Iit_minorder        sql.NullInt64
	Iit_active          sql.NullInt64
	Iit_taxable         sql.NullInt64
	Iit_default         sql.NullInt64
	Iit_buy             sql.NullInt64
	Iit_sell            sql.NullInt64
	Iit_inventory       sql.NullInt64
	Iit_allowdc         sql.NullInt64
	Iit_taxincl         sql.NullInt64
	Iit_reward          sql.NullInt64
	Iit_needref         sql.NullInt64
	Iit_temporary       sql.NullInt64
	Iit_fixedpromo      sql.NullInt64
	Iit_price_base      sql.NullFloat64
	Iit_price_part      sql.NullFloat64
	Iit_cost_avg        sql.NullFloat64
	Iit_cost_std        sql.NullFloat64
	Iit_cost_rtrn       sql.NullFloat64
	Iit_weight          sql.NullFloat64
	Iit_volume          sql.NullFloat64
	Iit_sale_price      sql.NullFloat64
	Iit_vendor_price    sql.NullFloat64
	Iit_sale_price_part sql.NullFloat64
	Iit_wprc_base       sql.NullFloat64
	Iit_wprc_part       sql.NullFloat64
	Iit_dclimit         sql.NullFloat64
	Iit_qdiscrate       sql.NullFloat64
	Iit_taxcut          sql.NullFloat64
	Iit_pointmul        sql.NullFloat64
	Iit_cost_last       sql.NullFloat64
	Iit_taxinprice      sql.NullFloat64
	Iit_exemlevel       sql.NullFloat64
	Iit_sale_begin      pq.NullTime
	Iit_sale_end        pq.NullTime
	Iit_created         pq.NullTime
	Iit_modified        pq.NullTime
}

func (iit Iit) NewItem() Iit {
	iit.Iit_code = ""
	iit.Iit_ict_code = sql.NullString{"", false}
	iit.Iit_pvd_code = sql.NullString{"", false}
	iit.Iit_ium_code = sql.NullString{"", false}
	iit.Iit_ity_code = sql.NullString{"", false}
	iit.Iit_ium_part = sql.NullString{"", false}
	iit.Iit_iwh_code = sql.NullString{"", false}
	iit.Iit_ytx_code = sql.NullString{"", false}
	iit.Iit_ocl_code = sql.NullString{"", false}
	iit.Iit_ycp_code = sql.NullString{"", false}
	iit.Iit_opr_code = sql.NullString{"", false}
	iit.Iit_oct_code = sql.NullString{"", false}
	iit.Iit_ist_code = sql.NullString{"", false}
	iit.Iit_ean13 = sql.NullString{"", false}
	iit.Iit_ean13a = sql.NullString{"", false}
	iit.Iit_name = sql.NullString{"", false}
	iit.Iit_namealt = sql.NullString{"", false}
	iit.Iit_vendor_sku = sql.NullString{"", false}
	iit.Iit_package = sql.NullString{"", false}
	iit.Iit_qty_base = sql.NullInt64{1, false}
	iit.Iit_qty_oh = sql.NullInt64{0, false}
	iit.Iit_pqty_oh = sql.NullInt64{0, false}
	iit.Iit_qty_bk = sql.NullInt64{0, false}
	iit.Iit_pqty_bk = sql.NullInt64{0, false}
	iit.Iit_qty_po = sql.NullInt64{0, false}
	iit.Iit_pqty_po = sql.NullInt64{0, false}
	iit.Iit_qty_it = sql.NullInt64{0, false}
	iit.Iit_pqty_it = sql.NullInt64{0, false}
	iit.Iit_qty_al = sql.NullInt64{0, false}
	iit.Iit_pqty_al = sql.NullInt64{0, false}
	iit.Iit_initqty = sql.NullInt64{0, false}
	iit.Iit_qdiscqty = sql.NullInt64{0, false}
	iit.Iit_point = sql.NullInt64{0, false}
	iit.Iit_leaddays = sql.NullInt64{0, false}
	iit.Iit_minlevel = sql.NullInt64{0, false}
	iit.Iit_minorder = sql.NullInt64{0, false}
	iit.Iit_active = sql.NullInt64{1, false}
	iit.Iit_taxable = sql.NullInt64{1, false}
	iit.Iit_default = sql.NullInt64{0, false}
	iit.Iit_buy = sql.NullInt64{1, false}
	iit.Iit_sell = sql.NullInt64{1, false}
	iit.Iit_inventory = sql.NullInt64{1, false}
	iit.Iit_allowdc = sql.NullInt64{0, false}
	iit.Iit_taxincl = sql.NullInt64{0, false}
	iit.Iit_reward = sql.NullInt64{0, false}
	iit.Iit_needref = sql.NullInt64{0, false}
	iit.Iit_temporary = sql.NullInt64{0, false}
	iit.Iit_fixedpromo = sql.NullInt64{0, false}
	iit.Iit_price_base = sql.NullFloat64{0.0, false}
	iit.Iit_price_part = sql.NullFloat64{0.0, false}
	iit.Iit_cost_avg = sql.NullFloat64{0.0, false}
	iit.Iit_cost_std = sql.NullFloat64{0.0, false}
	iit.Iit_cost_rtrn = sql.NullFloat64{0.0, false}
	iit.Iit_weight = sql.NullFloat64{0.0, false}
	iit.Iit_volume = sql.NullFloat64{0.0, false}
	iit.Iit_sale_price = sql.NullFloat64{0.0, false}
	iit.Iit_vendor_price = sql.NullFloat64{0.0, false}
	iit.Iit_sale_price_part = sql.NullFloat64{0.0, false}
	iit.Iit_wprc_base = sql.NullFloat64{0.0, false}
	iit.Iit_wprc_part = sql.NullFloat64{0.0, false}
	iit.Iit_dclimit = sql.NullFloat64{0.0, false}
	iit.Iit_qdiscrate = sql.NullFloat64{0.0, false}
	iit.Iit_taxcut = sql.NullFloat64{0.0, false}
	iit.Iit_pointmul = sql.NullFloat64{1.0, false}
	iit.Iit_cost_last = sql.NullFloat64{0.0, false}
	iit.Iit_taxinprice = sql.NullFloat64{0.0, false}
	iit.Iit_exemlevel = sql.NullFloat64{0.0, false}
	iit.Iit_sale_begin = pq.NullTime{time.Time{}, false}
	iit.Iit_sale_end = pq.NullTime{time.Time{}, false}
	iit.Iit_created = pq.NullTime{time.Now(), false}
	iit.Iit_modified = pq.NullTime{time.Now(), false}
	return iit
}

func (t Iit) GetItems(filter string) []Iit {
	var query string
	flist := strings.Split(filter, " ")
	if len(flist) > 0 {
		wherestr := "1=1"
		for _, ft := range flist {
			wft := strings.Replace(ft, "'", "\\'", -1)
			wherestr = wherestr + fmt.Sprintf(" AND (iit_code ilike '%%%s%%' OR iit_name ilike '%%%s%%' OR iit_name ilike '%%%s%%') ", wft, wft, wft)
		}
		query = fmt.Sprintf("SELECT iit_code, iit_ict_code, iit_pvd_code, iit_ium_code, iit_ity_code, iit_ium_part, iit_iwh_code, iit_ytx_code, iit_ocl_code, iit_ycp_code, iit_opr_code, iit_oct_code, iit_ist_code, iit_ean13, iit_ean13a, iit_name, iit_namealt, iit_vendor_sku, iit_package, iit_qty_base, iit_qty_oh, iit_pqty_oh, iit_qty_bk, iit_pqty_bk, iit_qty_po, iit_pqty_po, iit_qty_it, iit_pqty_it, iit_qty_al, iit_pqty_al, iit_initqty, iit_qdiscqty, iit_point, iit_leaddays, iit_minlevel, iit_minorder, iit_active, iit_taxable, iit_default, iit_buy, iit_sell, iit_inventory, iit_allowdc, iit_taxincl, iit_reward, iit_needref, iit_temporary, iit_fixedpromo, iit_price_base, iit_price_part, iit_cost_avg, iit_cost_std, iit_cost_rtrn, iit_weight, iit_volume, iit_sale_price, iit_vendor_price, iit_sale_price_part, iit_wprc_base, iit_wprc_part, iit_dclimit, iit_qdiscrate, iit_taxcut, iit_pointmul, iit_cost_last, iit_taxinprice, iit_exemlevel, iit_sale_begin, iit_sale_end, iit_created, iit_modified FROM iit_items WHERE %s ORDER BY iit_code ", wherestr)
	} else {
		query = "SELECT iit_code, iit_ict_code, iit_pvd_code, iit_ium_code, iit_ity_code, iit_ium_part, iit_iwh_code, iit_ytx_code, iit_ocl_code, iit_ycp_code, iit_opr_code, iit_oct_code, iit_ist_code, iit_ean13, iit_ean13a, iit_name, iit_namealt, iit_vendor_sku, iit_package, iit_qty_base, iit_qty_oh, iit_pqty_oh, iit_qty_bk, iit_pqty_bk, iit_qty_po, iit_pqty_po, iit_qty_it, iit_pqty_it, iit_qty_al, iit_pqty_al, iit_initqty, iit_qdiscqty, iit_point, iit_leaddays, iit_minlevel, iit_minorder, iit_active, iit_taxable, iit_default, iit_buy, iit_sell, iit_inventory, iit_allowdc, iit_taxincl, iit_reward, iit_needref, iit_temporary, iit_fixedpromo, iit_price_base, iit_price_part, iit_cost_avg, iit_cost_std, iit_cost_rtrn, iit_weight, iit_volume, iit_sale_price, iit_vendor_price, iit_sale_price_part, iit_wprc_base, iit_wprc_part, iit_dclimit, iit_qdiscrate, iit_taxcut, iit_pointmul, iit_cost_last, iit_taxinprice, iit_exemlevel, iit_sale_begin, iit_sale_end, iit_created, iit_modified FROM iit_items ORDER BY iit_code "

	}
	//fmt.Println(query)
	rows, err := db.Query(query)
	//fmt.Println(err)
	defer rows.Close()

	iits := make([]Iit, 0)
	for rows.Next() {
		iit := t.NewItem()
		err := rows.Scan(&iit.Iit_code, &iit.Iit_ict_code, &iit.Iit_pvd_code, &iit.Iit_ium_code, &iit.Iit_ity_code, &iit.Iit_ium_part, &iit.Iit_iwh_code, &iit.Iit_ytx_code, &iit.Iit_ocl_code, &iit.Iit_ycp_code, &iit.Iit_opr_code, &iit.Iit_oct_code, &iit.Iit_ist_code, &iit.Iit_ean13, &iit.Iit_ean13a, &iit.Iit_name, &iit.Iit_namealt, &iit.Iit_vendor_sku, &iit.Iit_package, &iit.Iit_qty_base, &iit.Iit_qty_oh, &iit.Iit_pqty_oh, &iit.Iit_qty_bk, &iit.Iit_pqty_bk, &iit.Iit_qty_po, &iit.Iit_pqty_po, &iit.Iit_qty_it, &iit.Iit_pqty_it, &iit.Iit_qty_al, &iit.Iit_pqty_al, &iit.Iit_initqty, &iit.Iit_qdiscqty, &iit.Iit_point, &iit.Iit_leaddays, &iit.Iit_minlevel, &iit.Iit_minorder, &iit.Iit_active, &iit.Iit_taxable, &iit.Iit_default, &iit.Iit_buy, &iit.Iit_sell, &iit.Iit_inventory, &iit.Iit_allowdc, &iit.Iit_taxincl, &iit.Iit_reward, &iit.Iit_needref, &iit.Iit_temporary, &iit.Iit_fixedpromo, &iit.Iit_price_base, &iit.Iit_price_part, &iit.Iit_cost_avg, &iit.Iit_cost_std, &iit.Iit_cost_rtrn, &iit.Iit_weight, &iit.Iit_volume, &iit.Iit_sale_price, &iit.Iit_vendor_price, &iit.Iit_sale_price_part, &iit.Iit_wprc_base, &iit.Iit_wprc_part, &iit.Iit_dclimit, &iit.Iit_qdiscrate, &iit.Iit_taxcut, &iit.Iit_pointmul, &iit.Iit_cost_last, &iit.Iit_taxinprice, &iit.Iit_exemlevel, &iit.Iit_sale_begin, &iit.Iit_sale_end, &iit.Iit_created, &iit.Iit_modified)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(iit.Iit_code)
		iits = append(iits, iit)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(iits)
	return iits
}

func (t Iit) GetItem(filter string) Iit {
	query := "SELECT iit_code, iit_ict_code, iit_pvd_code, iit_ium_code, iit_ity_code, iit_ium_part, iit_iwh_code, iit_ytx_code, iit_ocl_code, iit_ycp_code, iit_opr_code, iit_oct_code, iit_ist_code, iit_ean13, iit_ean13a, iit_name, iit_namealt, iit_vendor_sku, iit_package, iit_qty_base, iit_qty_oh, iit_pqty_oh, iit_qty_bk, iit_pqty_bk, iit_qty_po, iit_pqty_po, iit_qty_it, iit_pqty_it, iit_qty_al, iit_pqty_al, iit_initqty, iit_qdiscqty, iit_point, iit_leaddays, iit_minlevel, iit_minorder, iit_active, iit_taxable, iit_default, iit_buy, iit_sell, iit_inventory, iit_allowdc, iit_taxincl, iit_reward, iit_needref, iit_temporary, iit_fixedpromo, iit_price_base, iit_price_part, iit_cost_avg, iit_cost_std, iit_cost_rtrn, iit_weight, iit_volume, iit_sale_price, iit_vendor_price, iit_sale_price_part, iit_wprc_base, iit_wprc_part, iit_dclimit, iit_qdiscrate, iit_taxcut, iit_pointmul, iit_cost_last, iit_taxinprice, iit_exemlevel, iit_sale_begin, iit_sale_end, iit_created, iit_modified FROM iit_items WHERE iit_code = $1 ORDER BY iit_code "
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("filter: ", filter)
	rows, err := stmt.Query(filter)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(err)
	defer rows.Close()

	iit := t.NewItem()
	if rows.Next() {
		err := rows.Scan(&iit.Iit_code, &iit.Iit_ict_code, &iit.Iit_pvd_code, &iit.Iit_ium_code, &iit.Iit_ity_code, &iit.Iit_ium_part, &iit.Iit_iwh_code, &iit.Iit_ytx_code, &iit.Iit_ocl_code, &iit.Iit_ycp_code, &iit.Iit_opr_code, &iit.Iit_oct_code, &iit.Iit_ist_code, &iit.Iit_ean13, &iit.Iit_ean13a, &iit.Iit_name, &iit.Iit_namealt, &iit.Iit_vendor_sku, &iit.Iit_package, &iit.Iit_qty_base, &iit.Iit_qty_oh, &iit.Iit_pqty_oh, &iit.Iit_qty_bk, &iit.Iit_pqty_bk, &iit.Iit_qty_po, &iit.Iit_pqty_po, &iit.Iit_qty_it, &iit.Iit_pqty_it, &iit.Iit_qty_al, &iit.Iit_pqty_al, &iit.Iit_initqty, &iit.Iit_qdiscqty, &iit.Iit_point, &iit.Iit_leaddays, &iit.Iit_minlevel, &iit.Iit_minorder, &iit.Iit_active, &iit.Iit_taxable, &iit.Iit_default, &iit.Iit_buy, &iit.Iit_sell, &iit.Iit_inventory, &iit.Iit_allowdc, &iit.Iit_taxincl, &iit.Iit_reward, &iit.Iit_needref, &iit.Iit_temporary, &iit.Iit_fixedpromo, &iit.Iit_price_base, &iit.Iit_price_part, &iit.Iit_cost_avg, &iit.Iit_cost_std, &iit.Iit_cost_rtrn, &iit.Iit_weight, &iit.Iit_volume, &iit.Iit_sale_price, &iit.Iit_vendor_price, &iit.Iit_sale_price_part, &iit.Iit_wprc_base, &iit.Iit_wprc_part, &iit.Iit_dclimit, &iit.Iit_qdiscrate, &iit.Iit_taxcut, &iit.Iit_pointmul, &iit.Iit_cost_last, &iit.Iit_taxinprice, &iit.Iit_exemlevel, &iit.Iit_sale_begin, &iit.Iit_sale_end, &iit.Iit_created, &iit.Iit_modified)
		if err != nil {
			fmt.Println(err)
		}
	}
	//fmt.Println(query)
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(iits)
	return iit
}

func (t Iit) GetAlteredItems(filter string) []Iit {
	var query string
	flist := strings.Split(filter, " ")
	if len(flist) > 0 {
		wherestr := "1=1"
		for _, ft := range flist {
			wft := strings.Replace(ft, "'", "\\'", -1)
			wherestr = wherestr + fmt.Sprintf(" AND (iit_code ilike '%%%s%%' OR iit_name ilike '%%%s%%' OR iit_name ilike '%%%s%%') ", wft, wft, wft)
		}
		query = fmt.Sprintf("SELECT iit_code, iit_ict_code, iit_pvd_code, iit_ium_code, iit_ity_code, iit_ium_part, iit_iwh_code, iit_ytx_code, iit_ocl_code, iit_ycp_code, iit_opr_code, iit_oct_code, iit_ist_code, iit_ean13, iit_ean13a, iit_name, iit_namealt, iit_vendor_sku, iit_package, iit_qty_base, iit_qty_oh, iit_pqty_oh, iit_qty_bk, iit_pqty_bk, iit_qty_po, iit_pqty_po, iit_qty_it, iit_pqty_it, iit_qty_al, iit_pqty_al, iit_initqty, iit_qdiscqty, iit_point, iit_leaddays, iit_minlevel, iit_minorder, iit_active, iit_taxable, iit_default, iit_buy, iit_sell, iit_inventory, iit_allowdc, iit_taxincl, iit_reward, iit_needref, iit_temporary, iit_fixedpromo, iit_price_base, iit_price_part, iit_cost_avg, iit_cost_std, iit_cost_rtrn, iit_weight, iit_volume, iit_sale_price, iit_vendor_price, iit_sale_price_part, iit_wprc_base, iit_wprc_part, iit_dclimit, iit_qdiscrate, iit_taxcut, iit_pointmul, iit_cost_last, iit_taxinprice, iit_exemlevel, iit_sale_begin, iit_sale_end, iit_created, iit_modified FROM iit_items WHERE %s ORDER BY iit_code ", wherestr)
	} else {
		query = "SELECT iit_code, iit_ict_code, iit_pvd_code, iit_ium_code, iit_ity_code, iit_ium_part, iit_iwh_code, iit_ytx_code, iit_ocl_code, iit_ycp_code, iit_opr_code, iit_oct_code, iit_ist_code, iit_ean13, iit_ean13a, iit_name, iit_namealt, iit_vendor_sku, iit_package, iit_qty_base, iit_qty_oh, iit_pqty_oh, iit_qty_bk, iit_pqty_bk, iit_qty_po, iit_pqty_po, iit_qty_it, iit_pqty_it, iit_qty_al, iit_pqty_al, iit_initqty, iit_qdiscqty, iit_point, iit_leaddays, iit_minlevel, iit_minorder, iit_active, iit_taxable, iit_default, iit_buy, iit_sell, iit_inventory, iit_allowdc, iit_taxincl, iit_reward, iit_needref, iit_temporary, iit_fixedpromo, iit_price_base, iit_price_part, iit_cost_avg, iit_cost_std, iit_cost_rtrn, iit_weight, iit_volume, iit_sale_price, iit_vendor_price, iit_sale_price_part, iit_wprc_base, iit_wprc_part, iit_dclimit, iit_qdiscrate, iit_taxcut, iit_pointmul, iit_cost_last, iit_taxinprice, iit_exemlevel, iit_sale_begin, iit_sale_end, iit_created, iit_modified FROM iit_items ORDER BY iit_code "

	}
	//fmt.Println(query)
	rows, err := db.Query(query)
	//fmt.Println(err)
	defer rows.Close()

	iits := make([]Iit, 0)
	for rows.Next() {
		iit := t.NewItem()
		err := rows.Scan(&iit.Iit_code, &iit.Iit_ict_code, &iit.Iit_pvd_code, &iit.Iit_ium_code, &iit.Iit_ity_code, &iit.Iit_ium_part, &iit.Iit_iwh_code, &iit.Iit_ytx_code, &iit.Iit_ocl_code, &iit.Iit_ycp_code, &iit.Iit_opr_code, &iit.Iit_oct_code, &iit.Iit_ist_code, &iit.Iit_ean13, &iit.Iit_ean13a, &iit.Iit_name, &iit.Iit_namealt, &iit.Iit_vendor_sku, &iit.Iit_package, &iit.Iit_qty_base, &iit.Iit_qty_oh, &iit.Iit_pqty_oh, &iit.Iit_qty_bk, &iit.Iit_pqty_bk, &iit.Iit_qty_po, &iit.Iit_pqty_po, &iit.Iit_qty_it, &iit.Iit_pqty_it, &iit.Iit_qty_al, &iit.Iit_pqty_al, &iit.Iit_initqty, &iit.Iit_qdiscqty, &iit.Iit_point, &iit.Iit_leaddays, &iit.Iit_minlevel, &iit.Iit_minorder, &iit.Iit_active, &iit.Iit_taxable, &iit.Iit_default, &iit.Iit_buy, &iit.Iit_sell, &iit.Iit_inventory, &iit.Iit_allowdc, &iit.Iit_taxincl, &iit.Iit_reward, &iit.Iit_needref, &iit.Iit_temporary, &iit.Iit_fixedpromo, &iit.Iit_price_base, &iit.Iit_price_part, &iit.Iit_cost_avg, &iit.Iit_cost_std, &iit.Iit_cost_rtrn, &iit.Iit_weight, &iit.Iit_volume, &iit.Iit_sale_price, &iit.Iit_vendor_price, &iit.Iit_sale_price_part, &iit.Iit_wprc_base, &iit.Iit_wprc_part, &iit.Iit_dclimit, &iit.Iit_qdiscrate, &iit.Iit_taxcut, &iit.Iit_pointmul, &iit.Iit_cost_last, &iit.Iit_taxinprice, &iit.Iit_exemlevel, &iit.Iit_sale_begin, &iit.Iit_sale_end, &iit.Iit_created, &iit.Iit_modified)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(iit.Iit_code)
		iits = append(iits, iit)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(iits)
	return iits
}
