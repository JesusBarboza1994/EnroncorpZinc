package config
import(
	"testing"
)

func BenchmarkLoopUsers(b *testing.B) {
	// Establece el número de iteraciones para el benchmark
	b.N = 5

	// Ejecuta el benchmark
	for i := 0; i < b.N; i++ {
		LoopUsers("../../enron_mail_20110402/maildir")
	}
}