package queuesim

type Rule interface {
	Include(val int) bool
}

// Names of rules
const (
	WeightRuleName = "weight"
	HeightRuleName = "height"
)

type WeightRule []int

func NewWeightRule(vals ...int) Rule {
	return WeightRule(vals)
}

func (r WeightRule) Include(val int) bool {
	for _, v := range r {
		if v == val {
			return true
		}
	}

	return false
}

type HeightRule []int

func NewHeightRule(vals ...int) Rule {
	return HeightRule(vals)
}

func (r HeightRule) Include(val int) bool {
	for _, v := range r {
		if v == val {
			return true
		}
	}

	return false
}
