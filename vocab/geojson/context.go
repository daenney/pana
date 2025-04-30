package geojson

import (
	_ "embed"
)

//go:embed context.jsonld
var ContextDocument []byte

// IRI is the remote context IRI.
const IRI = "https://geojson.org/geojson-ld/geojson-context.jsonld"

// Namespace is the IRI prefix used for terms defined in this context that don't
// map to a different namespace.
const Namespace = "https://purl.org/geojson/vocab#"

const (
	// Bbox is a string.
	Bbox = Namespace + "bbox"
	// Coordinates is a string.
	Coordinates = Namespace + "coordinates"
	// Features is a string.
	Features = Namespace + "features"
	// Geometry is a string.
	Geometry = Namespace + "geometry"
	// Properties is a string.
	Properties = Namespace + "properties"
	// TypeFeature is a possible value for the type property.
	TypeFeature = Namespace + "Feature"
	// TypeFeatureCollection is a possible value for the type property.
	TypeFeatureCollection = Namespace + "FeatureCollection"
	// TypeGeometryCollection is a possible value for the type property.
	TypeGeometryCollection = Namespace + "GeometryCollection"
	// TypeLineString is a possible value for the type property.
	TypeLineString = Namespace + "LineString"
	// TypeMultiLineString is a possible value for the type property.
	TypeMultiLineString = Namespace + "MultiLineString"
	// TypeMultiPoint is a possible value for the type property.
	TypeMultiPoint = Namespace + "MultiPoint"
	// TypeMultiPolygon is a possible value for the type property.
	TypeMultiPolygon = Namespace + "MultiPolygon"
	// TypePoint is a possible value for the type property.
	TypePoint = Namespace + "Point"
	// TypePolygon is a possible value for the type property.
	TypePolygon = Namespace + "Polygon"
)
