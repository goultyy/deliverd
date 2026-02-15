/*
   This file contains definitions for:
   couriers, partners, classification
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
	err = db.QueryRow("SELECT * FROM deliverd.couriers WHERE courier_id = ?", courier_id).Scan(&courier.CourierID, &courier.NameFirst, &courier.NameMiddle, &courier.NameLast, &courier.AddressFirst, &courier.AddressSecond, &courier.AddressPostcode, &courier.AddressCounty, &courier.AddressCountry, &courier.Email, &courier.PhoneNumber, &courier.Status, &courier.DateBirth, &courier.DateJoined, &courier.DateLeft)
	return courier, err
}

// get all couriers
func get_all_couriers() ([]Courier, error) {
	db, err := get_db()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM deliverd.couriers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var couriers []Courier
	var courier Courier
	for rows.Next() {
		err = rows.Scan(&courier.CourierID, &courier.NameFirst, &courier.NameMiddle, &courier.NameLast, &courier.AddressFirst, &courier.AddressSecond, &courier.AddressPostcode, &courier.AddressCounty, &courier.AddressCountry, &courier.Email, &courier.PhoneNumber, &courier.Status, &courier.DateBirth, &courier.DateJoined, &courier.DateLeft)
		if err != nil {
			return couriers, err
		}
		couriers = append(couriers, courier)
	}
	return couriers, nil
}

// create a courier
func (courier Courier) new_courier() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO deliverd.couriers (name_first, name_middle, name_last, address_first, address_second, address_postcode, address_county, address_country, email, phone_number, status, date_birth, date_joined, date_left) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 0)", courier.NameFirst, courier.NameMiddle, courier.NameLast, courier.AddressFirst, courier.AddressSecond, courier.AddressPostcode, courier.AddressCounty, courier.AddressCountry, courier.Email, courier.PhoneNumber, courier.Status, courier.DateBirth, get_epoch())
	return err
}

// this will mark the courier as being on leave
func (courier Courier) mark_leave() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE deliverd.couriers SET status = ? WHERE courier_id = ?", COURIER_ON_LEAVE, courier.CourierID)
	change_key_status_by_owner(courier.CourierID, STATUS_INACTIVE)
	return err
}

// this will return the courier to being active (only if they have not left)
func (courier Courier) mark_active() error {
	if courier.Status != COURIER_ON_LEAVE {
		return errors.New("courier is not presently on leave and therefore cannot be marked as active")
	}
	change_key_status_by_owner(courier.CourierID, STATUS_ACTIVE)
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE deliverd.couriers SET status = ? WHERE courier_id = ?", COURIER_ACTIVE, courier.CourierID)
	return err
}

// this will mark the courier as having left (cannot be undone)
func (courier Courier) mark_left() error {
	db, err := get_db()
	change_key_status_by_owner(courier.CourierID, STATUS_INACTIVE)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE deliverd.couriers SET status = ?, date_left = ? WHERE courier_id = ?", COURIER_INACTIVE, get_epoch(), courier.CourierID)
	return err
}

// get partner by unqiue id
func get_partner_by_id(partner_id int) (Partner, error) {
	db, err := get_db()
	if err != nil {
		return Partner{}, err
	}
	var partner Partner
	err = db.QueryRow("SELECT * FROM deliverd.partners WHERE partners_id = ?", partner_id).Scan(&partner.PartnersID, &partner.CompanyName, &partner.CompanyID, &partner.AddressFirst, &partner.AddressSecond, &partner.AddressPostcode, &partner.AddressCounty, &partner.AddressCountry, &partner.Status, &partner.DateCreation, &partner.RepNameFirst, &partner.RepNameLast, &partner.RepEmail, &partner.RepPhone)
	return partner, err
}

// change status of a partner
func change_partner_status(partner_id int, status PartnerStatus) error {
	db, err := get_db()
	if err != nil {
		return err
	}
	partner, err := get_partner_by_id(partner_id)
	if err != nil {
		return err
	}
	if partner.Status == PARTNER_ACTIVE && status != PARTNER_INACTIVE || partner.Status == PARTNER_INACTIVE && status != PARTNER_ACTIVE {
		return errors.New("invalid status change")
	}

	if status == PARTNER_ACTIVE {
		change_key_status_by_owner(partner_id, STATUS_ACTIVE)
	} else {
		change_key_status_by_owner(partner_id, STATUS_INACTIVE)
	}
	_, err = db.Exec("UPDATE deliverd.partners SET status = ? WHERE partners_id = ?", status, partner.PartnersID)
	return err
}

// get all partners
func get_all_partners() ([]Partner, error) {
	db, err := get_db()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM deliverd.partners")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var partners []Partner
	var partner Partner
	for rows.Next() {
		err = rows.Scan(&partner.PartnersID, &partner.CompanyName, &partner.CompanyID, &partner.AddressFirst, &partner.AddressSecond, &partner.AddressPostcode, &partner.AddressCounty, &partner.AddressCountry, &partner.Status, &partner.DateCreation, &partner.RepNameFirst, &partner.RepNameLast, &partner.RepEmail, &partner.RepPhone)
		if err != nil {
			return nil, err
		}
		partners = append(partners, partner)
	}
	return partners, nil
}

// create a partner
func (partner Partner) new_partner() error {
	db, err := get_db()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO deliverd.partners (company_name, company_id, address_first, address_second, address_postcode, address_county, address_country, status, date_creation, rep_name_first, rep_name_last, rep_email, rep_phone) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", partner.CompanyName, partner.CompanyID, partner.AddressFirst, partner.AddressSecond, partner.AddressPostcode, partner.AddressCounty, partner.AddressCountry, PARTNER_ACTIVE, get_epoch(), partner.RepNameFirst, partner.RepNameLast, partner.RepEmail, partner.RepPhone)
	return err
}

// get classification by an id
func get_classification_by_id(classification_id int) (Classification, error) {
	db, err := get_db()
	if err != nil {
		return Classification{}, err
	}
	var classification Classification
	err = db.QueryRow("SELECT * FROM deliverd.classification WHERE classification_id = ?", classification_id).Scan(&classification.ClassificationID, &classification.Name, &classification.Description, &classification.ExpectedTime, &classification.RequiredTime, &classification.PriceGBP)
	return classification, err
}

// return all classifications
func get_all_classifications() ([]Classification, error) {
	db, err := get_db()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM deliverd.classification")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classifications []Classification
	var classification Classification
	for rows.Next() {
		err = rows.Scan(&classification.ClassificationID, &classification.Name, &classification.Description, &classification.ExpectedTime, &classification.RequiredTime, &classification.PriceGBP)
		if err != nil {
			return nil, err
		}
		classifications = append(classifications, classification)
	}
	return classifications, nil
}
