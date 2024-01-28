package sailboat

type Sailboat struct {
	Brand  string
	Color  string
	Length int
}

type SailboatBuilder struct {
	boat *Sailboat
}

func NewSailboatBuilder() *SailboatBuilder {
	return &SailboatBuilder{boat: &Sailboat{
		Brand:  "DefaultBrand",
		Color:  "DefaultColor",
		Length: 999,
	}}
}

func (sb *SailboatBuilder) Build() *Sailboat {
	return sb.boat
}

func (sb *SailboatBuilder) Brand(brand string) *SailboatBuilder {
	sb.boat.Brand = brand
	return sb
}

func (sb *SailboatBuilder) Color(color string) *SailboatBuilder {
	sb.boat.Color = color
	return sb
}

func (sb *SailboatBuilder) Length(length int) *SailboatBuilder {
	sb.boat.Length = length
	return sb
}