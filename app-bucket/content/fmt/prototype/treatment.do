clear
set seed 0
set obs 100


gen growthQ4_2025 = rnormal(0.01,1)
gen growthQ1_2026 = growthQ4_2025 + rnormal(0.01,1)
gen growthQ2_2026 = growthQ1_2026 + rnormal(0.01,1)

egen consensusQ4_2025 = mean(growthQ4_2025)
egen consensusQ1_2026 = mean(growthQ1_2026)
egen consensusQ2_2026 = mean(growthQ2_2026)

cumul growthQ4_2025, gen(grshareQ4_2025) equal
cumul growthQ1_2026, gen(grshareQ1_2026) equal
cumul growthQ2_2026, gen(grshareQ2_2026) equal

kdensity growthQ1_2026, xline(`=consensusQ1_2026[1]')  xline(`=growthQ1_2026[1]', lcolor(red))