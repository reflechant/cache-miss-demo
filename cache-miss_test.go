package cachemiss

import (
	"math/rand"
	"testing"
)

const k int = 100

func BenchmarkCPAOOP(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	banners := make([]AdBannerStat, b.N)
	for i := range banners {
		banners[i] = AdBannerStat{
			Impressions: rand.Uint64(),
			Spent:       rand.Uint64(),
			action:      rand.Intn(3),
		}
	}
	c := AdCampaignStat{Banners: banners}
	b.ResetTimer()
	// _ = c.CPM()
	var bs float64
	for j := 0; j < k; j++ {
		bs += c.CPA()
	}
	// }
}

func BenchmarkCPAOOPPtr(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	banners := make([]*AdBannerStat, b.N)
	for i := range banners {
		banners[i] = &AdBannerStat{
			Impressions: rand.Uint64(),
			Spent:       rand.Uint64(),
			action:      rand.Intn(3),
		}
	}
	c := AdCampaignStatPtr{Banners: banners}
	b.ResetTimer()
	// _ = c.CPM()
	var bs float64
	for j := 0; j < k; j++ {
		bs += c.CPA()
	}
	// }
}

func BenchmarkCPAOOPPtrRange(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	banners := make([]*AdBannerStat, b.N)
	for i := range banners {
		banners[i] = &AdBannerStat{
			Impressions: rand.Uint64(),
			Spent:       rand.Uint64(),
			action:      rand.Intn(3),
		}
	}
	c := AdCampaignStatPtr{Banners: banners}
	b.ResetTimer()
	// _ = c.CPM()
	var bs float64
	for j := 0; j < k; j++ {
		bs += c.CPArange()
	}
	// }
}

func BenchmarkCPAOOPIface(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	banners := make([]BannerStat, b.N)
	for i := range banners {
		banners[i] = &AdBannerStat{
			Impressions: rand.Uint64(),
			Spent:       rand.Uint64(),
			action:      rand.Intn(3),
		}
	}
	b.ResetTimer()
	// _ = c.CPM()
	var bs float64
	for j := 0; j < k; j++ {
		bs += CampaignCPAIface(banners)
	}
	// }
}

func BenchmarkCPADOD(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	impressions := make([]uint64, b.N)
	spent := make([]uint64, b.N)
	for i := 0; i < b.N; i++ {
		impressions[i] = rand.Uint64()
		spent[i] = rand.Uint64()
	}
	b.ResetTimer()
	// _ = CampaignCPM(impressions, spent)
	var bs float64
	for j := 0; j < k; j++ {
		bs += CampaignCPA(impressions, spent)
	}
	// }
}

func BenchmarkCPADODTwoLoops(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	impressions := make([]uint64, b.N)
	spent := make([]uint64, b.N)
	for i := 0; i < b.N; i++ {
		impressions[i] = rand.Uint64()
		spent[i] = rand.Uint64()
	}
	b.ResetTimer()
	// _ = CampaignCPM(impressions, spent)
	var bs float64
	for j := 0; j < k; j++ {
		bs += CampaignCPATwoLoops(impressions, spent)
	}
	// }
}
