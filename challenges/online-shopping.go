package challenges

type MahsulotTuri struct {
	Nomi string
	Narx float64
}

type Mijoz struct {
	Ismi     string
	Email    string
	Savatcha []MahsulotTuri
}

func (c *Mijoz) SavatchagaQosh(mahsulot MahsulotTuri) {
	c.Savatcha = append(c.Savatcha, mahsulot)
}

func (c *Mijoz) SavatchadanOchir(mahsulot MahsulotTuri) {
	yangiSavatcha := []MahsulotTuri{}

	for _, v := range c.Savatcha {
		if v == mahsulot {
			continue
		}
		yangiSavatcha = append(yangiSavatcha, v)
	}
	c.Savatcha = yangiSavatcha
}
