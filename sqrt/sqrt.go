package sqrt

import (
	"encoding/json"
	"math"
	"math/big"
	"net/http"
	"strconv"

	"github.com/jldec/learn-go-part2/counter"
)

// Sqrt algorithm started from https://tour.golang.org/flowcontrol/8
func Sqrt(x float64) (z float64, i int) {
	if x < 0 {
		return math.NaN(), 0
	}
	z = 1.0   // guess
	i = 1     // iterations
	s := true // delta sign
	f := 0    // delta sign changes
	for i < 1000 && f < 3 {
		znext := z - (z*z-x)/(2*z) // this overflows at z ~= 1e76
		if znext == z {
			return
		}
		delta := (znext - z) / z
		// fmt.Println("delta", delta)
		snext := (delta > 0)
		if snext != s {
			f++
		}
		z = znext
		s = snext
		i++
	}
	return
}

// HTTP request handler
// Input passed as parameter 'n'
// Returns JSON results
func Handler(cnt *counter.Counter) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cnt.Inc()

		input, err := strconv.ParseFloat(req.FormValue("n"), 64)
		if err != nil {
			http.Error(res, "Error parsing input n: "+err.Error(), 400)
			return
		}

		r := struct {
			Input, Sqrt, Math            float64
			NaNSqrt, NaNMath, SameResult bool
			Iterations                   int
			InputExp, SqrtExp, MathExp   int
		}{}

		r.Input = input
		r.InputExp = big.NewFloat(input).MantExp(nil)

		r.Sqrt, r.Iterations = Sqrt(input)
		if math.IsNaN(r.Sqrt) {
			r.NaNSqrt = true
			r.Sqrt = -1
		}
		r.SqrtExp = big.NewFloat(r.Sqrt).MantExp(nil)

		r.Math = math.Sqrt(input)
		if math.IsNaN(r.Math) {
			r.NaNMath = true
			r.Sqrt = -1
		}
		r.MathExp = big.NewFloat(r.Math).MantExp(nil)

		r.SameResult = r.Sqrt == r.Math

		enc := json.NewEncoder(res)
		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		enc.SetIndent("", "  ")
		err = enc.Encode(r)
		if err != nil {
			http.Error(res, "Error encoding result to JSON: "+err.Error(), 500)
		}
	})
}
