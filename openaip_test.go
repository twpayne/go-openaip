package openaip

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestDecodeAlt(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want Alt
	}{
		{`<ALT UNIT="FL">100</ALT>`, Alt{Unit: "FL", Value: 100}},
		{`<ALT UNIT="FL">75</ALT>`, Alt{Unit: "FL", Value: 75}},
	} {
		d := xml.NewDecoder(bytes.NewBufferString(tc.s))
		var got Alt
		if err := d.Decode(&got); err != nil {
			t.Errorf("decoding %#v returns %v, want nil", tc.s, err)
			continue
		}
		if got != tc.want {
			t.Errorf("decoding %#v yields %#v, want %#v", tc.s, got, tc.want)
		}
	}
}

func TestDecodeAltLimit(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want AltLimit
	}{
		{
			s:    `<ALTLIMIT_TOP REFERENCE="STD"><ALT UNIT="FL">100</ALT></ALTLIMIT_TOP>`,
			want: AltLimit{Reference: "STD", Value: Alt{Unit: "FL", Value: 100}},
		},
		{
			s:    `<ALTLIMIT_BOTTOM REFERENCE="STD"><ALT UNIT="FL">75</ALT></ALTLIMIT_BOTTOM>`,
			want: AltLimit{Reference: "STD", Value: Alt{Unit: "FL", Value: 75}},
		},
	} {
		d := xml.NewDecoder(bytes.NewBufferString(tc.s))
		var got AltLimit
		if err := d.Decode(&got); err != nil {
			t.Errorf("decoding %#v returns %v, want nil", tc.s, err)
			continue
		}
		if got != tc.want {
			t.Errorf("decoding %#v yields %#v, want %#v", tc.s, got, tc.want)
		}
	}
}

func TestDecodePolygon(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want Polygon
	}{
		{
			s: `<POLYGON>` +
				`9.6255555555556 48.5625, ` +
				`9.8477777777778 48.659444444444, ` +
				`9.8488590649036 48.671520456103, ` +
				`9.8491357659723 48.676179034769, ` +
				`9.8494899239756 48.688273884048, ` +
				`9.8494860905378 48.692936043918, ` +
				`9.8477777777778 48.705, ` +
				`9.9372222222222 48.714166666667, ` +
				`9.9380313248976 48.696291564689, ` +
				`9.937831067577 48.678408978203, ` +
				`9.9366228021006 48.660543713921, ` +
				`9.9344092727478 48.64272052819, ` +
				`9.9319444444444 48.630833333333, ` +
				`9.6391666666667 48.496388888889, ` +
				`9.6255555555556 48.5625` +
				`</POLYGON>`,
			want: Polygon{
				Coords: [][]float64{
					{9.6255555555556, 48.5625},
					{9.8477777777778, 48.659444444444},
					{9.8488590649036, 48.671520456103},
					{9.8491357659723, 48.676179034769},
					{9.8494899239756, 48.688273884048},
					{9.8494860905378, 48.692936043918},
					{9.8477777777778, 48.705},
					{9.9372222222222, 48.714166666667},
					{9.9380313248976, 48.696291564689},
					{9.937831067577, 48.678408978203},
					{9.9366228021006, 48.660543713921},
					{9.9344092727478, 48.64272052819},
					{9.9319444444444, 48.630833333333},
					{9.6391666666667, 48.496388888889},
					{9.6255555555556, 48.5625},
				},
			},
		},
	} {
		d := xml.NewDecoder(bytes.NewBufferString(tc.s))
		var got Polygon
		if err := d.Decode(&got); err != nil {
			t.Errorf("decoding %#v returns %v, want nil", tc.s, err)
			continue
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("decoding %#v yields %#v, want %#v", tc.s, got, tc.want)
		}
	}
}

func TestRead(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want OpenAIP
	}{
		{
			s: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>` +
				`<OPENAIP VERSION="367810a0f94887bf79cd9432d2a01142b0426795" DATAFORMAT="1.1">` +
				`<AIRSPACES>` +
				`<ASP CATEGORY="WAVE">` +
				`<VERSION>367810a0f94887bf79cd9432d2a01142b0426795</VERSION>` +
				`<ID>18024</ID>` +
				`<COUNTRY>DE</COUNTRY>` +
				`<NAME>ALB-OST</NAME>` +
				`<ALTLIMIT_TOP REFERENCE="STD">` +
				`<ALT UNIT="FL">100</ALT>` +
				`</ALTLIMIT_TOP>` +
				`<ALTLIMIT_BOTTOM REFERENCE="STD">` +
				`<ALT UNIT="FL">75</ALT>` +
				`</ALTLIMIT_BOTTOM>` +
				`<GEOMETRY>` +
				`<POLYGON>` +
				`9.6255555555556 48.5625, ` +
				`9.8477777777778 48.659444444444, ` +
				`9.8488590649036 48.671520456103, ` +
				`9.8491357659723 48.676179034769, ` +
				`9.8494899239756 48.688273884048, ` +
				`9.8494860905378 48.692936043918, ` +
				`9.8477777777778 48.705, ` +
				`9.9372222222222 48.714166666667, ` +
				`9.9380313248976 48.696291564689, ` +
				`9.937831067577 48.678408978203, ` +
				`9.9366228021006 48.660543713921, ` +
				`9.9344092727478 48.64272052819, ` +
				`9.9319444444444 48.630833333333, ` +
				`9.6391666666667 48.496388888889, ` +
				`9.6255555555556 48.5625` +
				`</POLYGON>` +
				`</GEOMETRY>` +
				`</ASP>` +
				`</AIRSPACES>` +
				`</OPENAIP>`,
			want: OpenAIP{
				Version:    "367810a0f94887bf79cd9432d2a01142b0426795",
				DataFormat: "1.1",
				Airspaces: []Airspace{
					{
						Category:       "WAVE",
						Version:        "367810a0f94887bf79cd9432d2a01142b0426795",
						Id:             18024,
						Country:        "DE",
						Name:           "ALB-OST",
						AltLimitTop:    AltLimit{Reference: "STD", Value: Alt{Unit: "FL", Value: 100}},
						AltLimitBottom: AltLimit{Reference: "STD", Value: Alt{Unit: "FL", Value: 75}},
						Polygons: []Polygon{
							{
								Coords: [][]float64{
									{9.6255555555556, 48.5625},
									{9.8477777777778, 48.659444444444},
									{9.8488590649036, 48.671520456103},
									{9.8491357659723, 48.676179034769},
									{9.8494899239756, 48.688273884048},
									{9.8494860905378, 48.692936043918},
									{9.8477777777778, 48.705},
									{9.9372222222222, 48.714166666667},
									{9.9380313248976, 48.696291564689},
									{9.937831067577, 48.678408978203},
									{9.9366228021006, 48.660543713921},
									{9.9344092727478, 48.64272052819},
									{9.9319444444444, 48.630833333333},
									{9.6391666666667, 48.496388888889},
									{9.6255555555556, 48.5625},
								},
							},
						},
					},
				},
			},
		},
	} {
		d := xml.NewDecoder(bytes.NewBufferString(tc.s))
		var got OpenAIP
		if err := d.Decode(&got); err != nil {
			t.Errorf("decoding %#v returns %v, want nil", tc.s, err)
			continue
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("decoding %#v yields %#v, want %#v", tc.s, got, tc.want)
		}
	}
}
