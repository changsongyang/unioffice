// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_TextBulletSize is a union type
type ST_TextBulletSize struct {
	ST_TextBulletSizePercent *string
	ST_TextBulletSizeDecimal *int32
}

func (m *ST_TextBulletSize) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextBulletSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_TextBulletSizePercent != nil {
		e.EncodeToken(xml.CharData(*m.ST_TextBulletSizePercent))
	}
	if m.ST_TextBulletSizeDecimal != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_TextBulletSizeDecimal)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_TextBulletSize) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_TextBulletSizePercent != nil {
		mems = append(mems, "ST_TextBulletSizePercent")
	}
	if m.ST_TextBulletSizeDecimal != nil {
		mems = append(mems, "ST_TextBulletSizeDecimal")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_TextBulletSize) String() string {
	if m.ST_TextBulletSizePercent != nil {
		return fmt.Sprintf("%v", *m.ST_TextBulletSizePercent)
	}
	if m.ST_TextBulletSizeDecimal != nil {
		return fmt.Sprintf("%v", *m.ST_TextBulletSizeDecimal)
	}
	return ""
}
