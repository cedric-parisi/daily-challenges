package axis

func fit1(x, y, a, b int) int {
	return (x / a) * (y / b)
}

func fit2(x, y, a, b int) int {
	n := fit1(x, y, a, b)
	r := fit1(x, y, b, a)
	if n >= r {
		return n
	}
	return r
}

func fit3(x, y, z, a, b, c int) int {
	res := []int{}
	res = append(res, (x/a)*(y/b)*(z/c))
	res = append(res, (x/a)*(y/c)*(z/b))
	res = append(res, (x/b)*(y/a)*(z/c))
	res = append(res, (x/c)*(y/a)*(z/b))
	res = append(res, (x/c)*(y/b)*(z/a))
	res = append(res, (x/b)*(y/c)*(z/a))
	max := res[0]
	for _, n := range res {
		if n > max {
			max = n
		}
	}
	return max
}

func fitn(crate, box []int) int {
	return 0
}
