/*
This file contains the code for the following types:
drops, packages, route, updates
*/
package deliverd

import "errors"

// Get drop by UniqueID
func get_drop_by_id(drop_id int) (Drops, error) {
	db, err := get_db()
	if err != nil {
		return Drops{}, err
	}
	var drop Drops
	err = db.QueryRow("SELECT * FROM deliverd.drops WHERE drop_id = ?", drop_id).Scan(&drop.DropID, &drop.PackageID, &drop.RouteID, &drop.Order)
	return drop, err
}

// Get drop by RouteID in order
func get_drops_by_route_id(route_id int) ([]Drops, error) {
	db, err := get_db()
	var drops []Drops
	if err != nil {
		return drops, err
	}
	var drop Drops
	rows, err := db.Query("SELECT * FROM deliverd.drops WHERE route_id = ? ORDER BY order ASC", route_id)
	if err != nil {
		return drops, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&drop.DropID, &drop.PackageID, &drop.RouteID, &drop.Order)
		if err != nil {
			return drops, err
		}
		drops = append(drops, drop)
	}
	return drops, err
}

// Short way to get ID
func get_package_id_by_drop_id(drop_id int) (int, error) {
	drop, err := get_drop_by_id(drop_id)
	if err != nil {
		return 0, err
	}
	return drop.PackageID, nil
}

// create drop from structure
func (drop Drops) new_drop() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO deliverd.drops (package_id, route_id, order) VALUES (?, ?, ?)", drop.PackageID, drop.RouteID, drop.Order)
	return err
}

// update package's status by package_id
func update_package_status(package_id int, status UpdateID) error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE deliverd.packages SET status = ? WHERE package_id = ?", status, package_id)
	return err
}

// update drop's relevant package using only drop id
func update_drop_status(drop_id, status UpdateID) error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE deliverd.packages SET status = ? WHERE package_id = (SELECT package_id FROM deliverd.drops WHERE drop_id = ?)", status, drop_id)
	return err
}

// get package by package id
func get_package_by_id(package_id int) (Packages, error) {
	db, err := get_db()
	if err != nil {
		return Packages{}, err
	}
	var pkg Packages
	err = db.QueryRow("SELECT * FROM deliverd.packages WHERE package_id = ?", package_id).Scan(&pkg.PackageID, &pkg.PartnerID, &pkg.Description, &pkg.DateCreated, &pkg.DestinationFirst, &pkg.DestinationSecond, &pkg.DestinationPostcode, &pkg.DestinationCounty, &pkg.DestinationCountry, &pkg.RecipientNameFirst, &pkg.RecipientNameLast, &pkg.RecipientEmail, &pkg.RecipientPhone, &pkg.Classification, &pkg.Status)
	return pkg, err
}

// get packages by status
func get_packages_by_status(status UpdateID) ([]Packages, error) {
	db, err := get_db()
	var pkgs []Packages
	if err != nil {
		return pkgs, err
	}
	var pkg Packages
	rows, err := db.Query("SELECT * FROM deliverd.packages WHERE status = ?", status)
	if err != nil {
		return pkgs, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&pkg.PackageID, &pkg.PartnerID, &pkg.Description, &pkg.DateCreated, &pkg.DestinationFirst, &pkg.DestinationSecond, &pkg.DestinationPostcode, &pkg.DestinationCounty, &pkg.DestinationCountry, &pkg.RecipientNameFirst, &pkg.RecipientNameLast, &pkg.RecipientEmail, &pkg.RecipientPhone, &pkg.Classification, &pkg.Status)
		if err != nil {
			return pkgs, err
		}
		pkgs = append(pkgs, pkg)
	}
	return pkgs, nil
}

// get packages by partner id
func get_packages_by_partner_id(partner_id int) ([]Packages, error) {
	db, err := get_db()
	var pkgs []Packages
	if err != nil {
		return pkgs, err
	}
	var pkg Packages
	rows, err := db.Query("SELECT * FROM deliverd.packages WHERE partner_id = ?", partner_id)
	if err != nil {
		return pkgs, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&pkg.PackageID, &pkg.PartnerID, &pkg.Description, &pkg.DateCreated, &pkg.DestinationFirst, &pkg.DestinationSecond, &pkg.DestinationPostcode, &pkg.DestinationCounty, &pkg.DestinationCountry, &pkg.RecipientNameFirst, &pkg.RecipientNameLast, &pkg.RecipientEmail, &pkg.RecipientPhone, &pkg.Classification, &pkg.Status)
		if err != nil {
			return pkgs, err
		}
		pkgs = append(pkgs, pkg)
	}
	return pkgs, nil
}

// create package
func (pkg Packages) new_package() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	partner, err := get_partner_by_id(pkg.PartnerID)
	if err != nil {
		return err
	}

	if partner.Status != PARTNER_ACTIVE {
		return errors.New("inactive partner cannot create package")
	}

	_, err = get_classification_by_id(pkg.Classification)
	if err != nil {
		return err
	}

	pkg.DateCreated = get_epoch()
	pkg.Status = UPDATE_INFORMED_SENDER
	_, err = db.Exec("INSERT INTO deliverd.packages (partner_id, description, date_created, destination_first, destination_second, destination_postcode, destination_county, destination_country, recipient_name_first, recipient_name_last, recipient_email, recipient_phone, classification, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", pkg.PartnerID, pkg.Description, pkg.DateCreated, pkg.DestinationFirst, pkg.DestinationSecond, pkg.DestinationPostcode, pkg.DestinationCounty, pkg.DestinationCountry, pkg.RecipientNameFirst, pkg.RecipientNameLast, pkg.RecipientEmail, pkg.RecipientPhone, pkg.Classification, pkg.Status)
	return err
}

// add package to route
func add_package_to_route(package_id int, route_id int) error {
	drops, err := get_drops_by_route_id(route_id)
	if err != nil {
		return err
	}
	order := len(drops) + 1
	drop := Drops{
		PackageID: package_id,
		RouteID:   route_id,
		Order:     order,
	}
	return drop.new_drop()
}

// swap order of two drops
func swap_drop_order(drop_id1, drop_id2 int) error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE deliverd.drops d1 JOIN deliverd.drops d2 ON d1.drop_id = ? AND d2.drop_id = ? SET d1.order = d2.order, d2.order = d1.order", drop_id1, drop_id2)
	return err
}

// new route (ignores Status)
func (route Route) new_route() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO deliverd.routes (courier_id, status, time_created, time_completed) VALUES (?, ?, ?, 0)", route.CourierID, ROUTE_ACTIVE, get_epoch())
	return err
}

func get_route_by_id(route_id int) (Route, error) {
	db, err := get_db()
	if err != nil {
		return Route{}, err
	}
	var route Route
	err = db.QueryRow("SELECT * FROM deliverd.routes WHERE route_id = ?", route_id).Scan(&route.RouteID, &route.CourierID, &route.Status, &route.TimeCreated, &route.TimeCompleted)
	return route, err
}

// create new drop update
func (update Update) new_update() error {
	_, err := get_package_by_id(update.PackageID)
	if err != nil {
		return err
	}
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO deliverd.updates (package_id, update_type, address_first, address_second, address_postcode, address_county, address_country, freeform, update_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", update.PackageID, update.UpdateType, update.AddressFirst, update.AddressSecond, update.AddressPostcode, update.AddressCounty, update.AddressCountry, update.Freeform, get_epoch())
	if err != nil {
		return err
	}
	err = update_package_status(update.PackageID, update.UpdateType)
	return err
}

// get all drop updates by package id
func get_updates_by_package_id(package_id int) ([]Update, error) {
	db, err := get_db()
	var updates []Update
	if err != nil {
		return updates, err
	}
	var update Update
	rows, err := db.Query("SELECT * FROM deliverd.updates WHERE package_id = ? ORDER BY update_date ASC", package_id)
	if err != nil {
		return updates, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&update.UpdateID, &update.PackageID, &update.UpdateType, &update.AddressFirst, &update.AddressSecond, &update.AddressPostcode, &update.AddressCounty, &update.AddressCountry, &update.Freeform, &update.UpdateDate)
		if err != nil {
			return updates, err
		}
		updates = append(updates, update)
	}
	return updates, nil
}
