package v1

import (
	pb "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func (s *service) GetContainerEvents(req *pb.GetEventsRequest, ces pb.RuntimeService_GetContainerEventsServer) error {

	for containerEvent := range s.server.ContainerEventsChan {
		if err := ces.Send(&containerEvent); err != nil {
			return err
		}
	}

	// Dummy code to generate a sample events stream
	// x := 1
	// for {
	// 	resp := pb.ContainerEventResponse{ContainerId: fmt.Sprintf("%v", x), ContainerEventType: pb.ContainerEventType_STOPPED}
	// 	if err := ces.Send(&resp); err != nil {
	// 		return err
	// 	}
	// 	x++
	// 	time.Sleep(time.Second * 3)
	// }

	return nil
}
