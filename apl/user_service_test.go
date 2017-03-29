package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
	"encoding/base64"
)

var (
	testUserId = "user-test-id"
	testUserFilter = "user"
)

func TestUserService_Create(t *testing.T) {
	aplSvs := apl.NewClient()

	data := []byte("transfloopiglorpsnarfle")
	passwd := base64.StdEncoding.EncodeToString(data)

	in := &apl.UserCreateInput{
		ID: testUserId,
		FirstName: "TestFirstName",
		LastName: "TestLastName",
		Email: "testing@example.com",
		Password: passwd,
		UserType: testUserFilter,
		WorkRole: "Unknown",
		RoleId: "ops-role-id",
	}

	out, _, err := aplSvs.Users.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	
	fmt.Println("New User ID:", out.Data)

}

func TestUserService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Users.List(nil)

	fmt.Println("count:", len(out))
	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("User found %d rows\n", rowCount)
	} else {
		t.Fatal("No User rows found")
	}

}

func TestUserService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.UserParams{
		UserType: testUserFilter,
	}
	out, _, err := aplSvc.Users.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No User rows found for filter", testUserFilter)
	}

	fmt.Printf("User filtered found %d rows for filter \"%s\"\n", rowCount, testUserFilter)

}

func TestUserService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.UserUpdateInput{FirstName: "UPDATED!"}
	out, _, err := aplSvc.Users.Update(testUserId, in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}

func TestUserService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Users.Get(testUserId)

	if err != nil {
		t.Fatal(err)
	}
	if out == (apl.User{}) {
		t.Fatal("No User found for ID", testUserId)
	}

	fmt.Println("User found for ID", testUserId)

}

func TestUserService_Delete(t *testing.T) {
	aplSvc := apl.NewClient()
	out, _, err := aplSvc.Users.Delete(testUserId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}