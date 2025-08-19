package tests

import (
	"Accounting/internal/entities"
	"log"
	"os"
	"reflect"
	"runtime/pprof"
	"testing"
	"unsafe"
)

/*
What is this file for?:
SetunexPortedfield - a universal function that allows you to change the private fields of the structure through Reflect + Unsafe.

TestAppuser_privatefields - Change the private fields Login and Password, check that Getuserlogin() returns the changed value.

TestAppuser_parallelheavy - a heavy test, chases in parallel (t.parallel()).

TestMemoryprofile - creates a million users and writes Memprofile.prof.
*/

func setUnexportedField(obj interface{}, fieldName string, value interface{}) {
	v := reflect.ValueOf(obj).Elem()
	f := v.FieldByName(fieldName)

	if !f.IsValid() {
		panic("field not found: " + fieldName)
	}

	reflect.NewAt(f.Type(), unsafe.Pointer(f.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf(value))
}

func TestAppUser_PrivateFields(t *testing.T) {
	t.Run("Modify login with reflect+unsafe", func(t *testing.T) {
		user := entities.RegisterUser("initial", "123")

		setUnexportedField(user, "login", "hacked_login")
		setUnexportedField(user, "password", "hacked_pass")

		got := user.GetUserLogin()
		want := "hacked_login"

		if got != want {
			t.Errorf("expected %s, got %s", want, got)
		}
	})
}

func TestAppUser_ParallelHeavy(t *testing.T) {
	t.Run("Parallel tests on terminalID", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < 10000; i++ {
			user := entities.RegisterUser("user", "pass")

			if user.GetUserLogin() != "user" {
				t.Errorf("unexpected login")
			}
		}
	})
}

func TestMemoryProfile(t *testing.T) {
	f, err := os.Create("memprofile.prof")
	if err != nil {
		t.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()

	for i := 0; i < 1_000_000; i++ {
		_ = entities.RegisterUser("bulkUser", "bulkPass")
	}

	if err := pprof.WriteHeapProfile(f); err != nil {
		t.Fatal("could not write memory profile: ", err)
	}
	log.Println("Memory profile written to memprofile.prof")
}
