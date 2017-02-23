package main

var query14 = `
SELECT
	100.00 * SUM(CASE
		WHEN p_type LIKE 'PROMO%'
			THEN l_extendedprice * (1 - l_discount)
		ELSE 0
	END) / SUM(l_extendedprice * (1 - l_discount)) AS promo_revenue
FROM
	lineitem,
	part
WHERE
	l_partkey = p_partkey
	AND l_shipdate >= DATE '1997-04-01'
	AND l_shipdate < DATE '1997-04-01' + INTERVAL '1' MONTH;
`
