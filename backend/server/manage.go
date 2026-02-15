/*
	This file contains definitions for:
	couriers, partners
*/

package deliverd

import "errors"

// get courier by it's Id
func get_courier_by_id(courier_id int) (Courier, error) {
	db, err := get_db()
	if err != nil {
		return Courier{}, err
	}
	var courier Courier
	err = db.QueryRow("SELECT * FROM couriers WHERE courier_id = ?", courier_id).Scan(&courier.CourierID, &courier.NameFirst, &courier.NameMiddle, &courier.NameLast, &courier.AddressFirst, &courier.AddressSecond, &courier.AddressPostcode, &courier.AddressCounty, &courier.AddressCountry, &courier.Email, &courier.PhoneNumber, &courier.Status, &courier.DateBirth, &courier.DateJoined, &courier.DateLeft)
	return courier, err
}

// create a courier
func (courier Courier) new_courier() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO couriers (name_first, name_middle, name_last, address_first, address_second, address_postcode, address_county, address_country, email, phone_number, status, date_birth, date_joined, date_left) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 0)", courier.NameFirst, courier.NameMiddle, courier.NameLast, courier.AddressFirst, courier.AddressSecond, courier.AddressPostcode, courier.AddressCounty, courier.AddressCountry, courier.Email, courier.PhoneNumber, courier.Status, courier.DateBirth, courier.DateJoined)
	return err
}

// this will mark the courier as being on leave
func (courier Courier) mark_leave() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE couriers SET status = ? WHERE courier_id = ?", COURIER_ON_LEAVE, courier.CourierID)
	return err
}

// this will return the courier to being active (only if they have not left)
func (courier Courier) mark_active() error {
	if courier.Status != COURIER_ON_LEAVE {
		return errors.New("courier is not presently on leave and therefore cannot be marked as active")
	}

	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE couriers SET status = ? WHERE courier_id = ?", COURIER_ACTIVE, courier.CourierID)
	return err
}

// this will mark the courier as having left (cannot be undone)
func (courier Courier) mark_left() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE couriers SET status = ?, date_left = ? WHERE courier_id = ?", COURIER_INACTIVE, get_epoch(), courier.CourierID)
	return err
}

// get partner by unqiue id
func get_partner_by_id(partner_id int) (Partner, error) {
	db, err := get_db()
	if err != nil {
		return Partner{}, err
	}
	var partner Partner
	err = db.QueryRow("SELECT * FROM partners WHERE partner_id = ?", partner_id).Scan(&partner.PartnersID, &partner.CompanyName, &partner.CompanyID, &partner.AddressFirst, &partner.AddressSecond, &partner.AddressPostcode, &partner.AddressCounty, &partner.AddressCountry, &partner.Status, &partner.DateCreation, &partner.RepNameFirst, &partner.RepNameLast, &partner.RepEmail, &partner.RepPhone)
	return partner, err
}

// create a partner
func (partner Partner) new_partner() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO partners (company_name, company_id, address_first, address_second, address_postcode, address_county, address_country, status, date_creation, rep_name_first, rep_name_last, rep_email, rep_phone) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", partner.CompanyName, partner.CompanyID, partner.AddressFirst, partner.AddressSecond, partner.AddressPostcode, partner.AddressCounty, partner.AddressCountry, partner.Status, partner.DateCreation, partner.RepNameFirst, partner.RepNameLast, partner.RepEmail, partner.RepPhone)
	return err
}
