package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
)

type User struct {
	Uid      string `json:"iud,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) createNewUser(c *dgo.Dgraph) error {
	ctx := context.Background()
	txn := c.NewTxn()
	defer txn.Discard(ctx)

	err := user.checkUserOrEmail(c)
	if err != nil {
		if IsValidationError(err) {
			return err
		}
		log.Fatal(err)
	}

	u, err := json.Marshal(&user)
	if err != nil {
		return err
	}

	mu := &api.Mutation{
		SetJson:   u,
		CommitNow: true,
	}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) checkUserOrEmail(c *dgo.Dgraph) error {
	ctx := context.Background()
	txn := c.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(`
	{
		user(func: eq(username, %q)) {
			username
		}
		email(func: eq(email, %q)) {
			email
		}
	}`, user.Username, user.Email)

	resp, err := txn.Query(ctx, q)
	if err != nil {
		return err
	}

	var data struct {
		Username []struct {
			Name string `json:"username"`
		} `json:"user"`
		Email []struct {
			Email string `json:"email"`
		} `json:"email"`
	}

	err = json.Unmarshal(resp.GetJson(), &data)
	if err != nil {
		return err
	}

	if len(data.Username) == 0 && len(data.Email) == 0 {
		return nil
	} else {
		for _, v := range data.Username {
			if v.Name == user.Username {
				return errUsernameExists
			}
		}
		for _, v := range data.Email {
			if v.Email == user.Email {
				return errEmailExists
			}
		}
	}

	return nil
}

func (user *User) login(c *dgo.Dgraph) error {
	ctx := context.Background()
	txn := c.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(`
	{
		user(func: eq(username, %q)) {
			checkpwd(password, %q)
		}
		email(func: eq(email, %q)) {
			checkpwd(password, %q)
		}
	}`, user.Username, user.Password, user.Email, user.Password)

	resp, err := txn.Query(ctx, q)
	if err != nil {
		return err
	}

	var data struct {
		Username []struct {
			CheckPwd bool `json:"checkpwd(password)"`
		} `json:"user"`
		Email []struct {
			CheckPwd bool `json:"checkpwd(password)"`
		} `json:"email"`
	}

	err = json.Unmarshal(resp.GetJson(), &data)
	if err != nil {
		return err
	}

	if len(data.Username) == 0 && len(data.Email) == 0 {
		return errCredentialsIncorrect
	} else {
		for _, v := range data.Username {
			if v.CheckPwd {
				return nil
			}
		}
		for _, v := range data.Email {
			if v.CheckPwd {
				return nil
			}
		}
	}

	return errCredentialsIncorrect
}
