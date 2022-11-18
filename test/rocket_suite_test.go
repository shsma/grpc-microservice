// +acceptance

package test

import (
	"context"
	rocket "github.com/shsma/grpc-microservice/proto/rocket/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

type RocketTestSuite struct {
	suite.Suite
}

func (s *RocketTestSuite) TestAddRocket() {
	s.T().Run("Adds a new rocket successfully", func(t *testing.T) {
		client := GetClient()
		res, err := client.AddRocket(context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "6",
					Name: "V1",
					Type: "Falcon-heavy",
				}})
		assert.NoError(t, err)
		assert.Equal(t, "6", res.Rocket.Id)
	})

	s.T().Run("validates the id in the new rocket is a number", func(t *testing.T) {
		client := GetClient()
		_, err := client.AddRocket(
			context.Background(),
			&rocket.AddRocketRequest{
				Rocket: &rocket.Rocket{
					Id:   "not-a-number-id",
					Name: "V1",
					Type: "Falcon Heavy",
				},
			},
		)
		assert.Error(s.T(), err)
		st := status.Convert(err)
		assert.Equal(s.T(), codes.InvalidArgument, st.Code())
	})
}

func TestRocketService(t *testing.T) {
	suite.Run(t, new(RocketTestSuite))
}
