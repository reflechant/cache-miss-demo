package cachemiss

// bad code, mainstream OOP stype

type AdBannerStat struct {
	Impressions       uint64
	UniqueImpressions uint64
	Clicks            uint64
	UniqueClicks      uint64
	Installs          uint64
	Spent             uint64
	Currency          string
	Name              string
	A                 string
	B                 string
	C                 string
	D                 string
	action            int
}

type AdCampaignStat struct {
	Banners []AdBannerStat
}

func (c *AdCampaignStat) CPA() float64 {
	var (
		total_impressions uint64
		total_spent       uint64
	)
	for i := range c.Banners {
		switch c.Banners[i].action {
		case 0:
			total_impressions += c.Banners[i].Impressions
		case 1:
			total_impressions += c.Banners[i].UniqueImpressions
		case 2:
			total_impressions += c.Banners[i].Clicks
		default:
			total_impressions += c.Banners[i].UniqueClicks
		}

		total_spent += c.Banners[i].Spent
	}
	if total_impressions == 0 {
		return 0.0
	}

	return float64(total_spent / total_impressions)
}

type AdCampaignStatPtr struct {
	Banners []*AdBannerStat
}

func (c *AdCampaignStatPtr) CPA() float64 {
	var (
		total_impressions uint64
		total_spent       uint64
	)
	for i := range c.Banners {
		switch c.Banners[i].action {
		case 0:
			total_impressions += c.Banners[i].Impressions
		case 1:
			total_impressions += c.Banners[i].UniqueImpressions
		case 2:
			total_impressions += c.Banners[i].Clicks
		default:
			total_impressions += c.Banners[i].UniqueClicks
		}
		total_spent += c.Banners[i].Spent
	}
	if total_impressions == 0 {
		return 0.0
	}

	return float64(total_spent / total_impressions)
}

// we use range-val instead if index-based loop

func (c *AdCampaignStatPtr) CPArange() float64 {
	var (
		total_impressions uint64
		total_spent       uint64
	)
	for _, b := range c.Banners {
		switch b.action {
		case 0:
			total_impressions += b.Impressions
		case 1:
			total_impressions += b.UniqueImpressions
		case 2:
			total_impressions += b.Clicks
		default:
			total_impressions += b.UniqueClicks
		}
		total_spent += b.Spent
	}
	if total_impressions == 0 {
		return 0.0
	}

	return float64(total_spent / total_impressions)
}

// even worse - create an interface

type BannerStat interface {
	Action() uint64
	Cost() uint64
}

func (bs *AdBannerStat) Cost() uint64 {
	return bs.Spent
}

func (bs *AdBannerStat) Action() uint64 {
	switch bs.action {
	case 0:
		return bs.Impressions
	case 1:
		return bs.UniqueImpressions
	case 2:
		return bs.Clicks
	default:
		return bs.UniqueClicks
	}
}

func CampaignCPAIface(banners []BannerStat) float64 {
	var (
		total_actions uint64
		total_spent   uint64
	)
	for _, b := range banners {
		total_actions += b.Action()
		total_spent += b.Cost()
	}
	if total_actions == 0 {
		return 0.0
	}

	return float64(total_spent / total_actions)
}

// good code, data-oriented style

func CampaignCPA(action []uint64, spent []uint64) float64 {
	var (
		total_impressions uint64
		total_spent       uint64
	)
	for i := 0; i < len(action); i++ {
		total_impressions += action[i]
		total_spent += spent[i]
	}
	if total_impressions == 0 {
		return 0.0
	}

	return float64(total_spent / total_impressions)
}

func CampaignCPATwoLoops(action []uint64, spent []uint64) float64 {
	var (
		total_impressions uint64
		total_spent       uint64
	)
	for i := 0; i < len(action); i++ {
		total_impressions += action[i]
	}
	for i := 0; i < len(spent); i++ {
		total_spent += spent[i]
	}
	if total_impressions == 0 {
		return 0.0
	}

	return float64(total_spent / total_impressions)
}
