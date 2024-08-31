package services

import (
	"context"

	"github.com/nrssi/kanbee/internal/db"
	// "github.com/nrssi/kanbee/internal/db/models"
	pb "github.com/nrssi/kanbee/internal/services/kanban"
	// kbp "github.com/nrssi/kanbee/internal/grpc/kanban/proto"
)

const TimeFormat string = "31/08/2024 - 12:59:13"

type KanbeeServiceServer struct {
	pb.UnimplementedKanbanServiceServer
}

func (s *KanbeeServiceServer) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.ProjectResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		project, err := db.CreateProject(req.Name, req.Description)
		if err != nil {
			return nil, err
		}
		return &pb.ProjectResponse{
			Project: &pb.Project{
				Id:          int32(project.ID),
				Name:        project.Name,
				Description: project.Description,
			},
		}, nil
	}
}

func (s *KanbeeServiceServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		task, err := db.CreateTask(req.Title, req.Description, uint(req.ProjectId), req.State)
		if err != nil {
			return nil, err
		}
		return &pb.TaskResponse{
			Task: &pb.Task{
				Id:          int32(task.ID),
				ProjectId:   int32(task.ProjectID),
				Title:       task.Title,
				Description: task.Description,
				State:       task.State,
				Deadline:    task.Deadline.Format(""),
			},
		}, nil
	}
}

func (s *KanbeeServiceServer) CreateUpdate(ctx context.Context, req *pb.CreateUpdateRequest) (*pb.UpdateResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		update, err := db.CreateUpdate(uint(req.TaskId), req.GetMessage(), req.GetFilename(), req.GetAttachmentData())
		if err != nil {
			return nil, err
		}
		return &pb.UpdateResponse{
			Update: &pb.Update{
				Id:             int32(update.ID),
				TaskId:         int32(update.TaskID),
				Message:        update.Message,
				AttachmentData: update.AttachmentData,
				Filename:       update.Filename,
				CreatedAt:      update.CreatedAt.Format(TimeFormat),
			},
		}, nil
	}
}

func (s *KanbeeServiceServer) UpdateProject(ctx context.Context, in *pb.UpdateProjectRequest) (*pb.ProjectResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		project, err := db.UpdateProject(uint(in.Id), in.Name, in.Description)
		if err != nil {
			return nil, err
		}
		return &pb.ProjectResponse{
			Project: &pb.Project{
				Id:          int32(project.ID),
				Name:        project.Name,
				Description: project.Description,
			},
		}, nil
	}
}

func (s *KanbeeServiceServer) GetProjectById(ctx context.Context, in *pb.GetByIdRequest) (*pb.ProjectResponse, error) {
	return nil, nil
}

func (s *KanbeeServiceServer) ListProjects(ctx context.Context, in *pb.Empty) (*pb.ListProjectsResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		results, err := db.ListProjects()
		if err != nil {
			return nil, err
		}
		var projects []*pb.Project
		for _, result := range results {
			project := &pb.Project{
				Id:          int32(result.ID),
				Name:        result.Name,
				Description: result.Name,
			}
			projects = append(projects, project)
		}
		return &pb.ListProjectsResponse{
			Projects: projects,
		}, nil
	}
}

func (s *KanbeeServiceServer) UpdateTask(ctx context.Context, in *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		task, err := db.UpdateTask(uint(in.Id), in.Title, in.Description, uint(in.ProjectId), in.State)
		if err != nil {
			return nil, err
		}
		return &pb.TaskResponse{
			Task: &pb.Task{
				Id:          int32(task.ID),
				ProjectId:   int32(task.ProjectID),
				Title:       task.Title,
				Description: task.Description,
				State:       task.State,
				Deadline:    task.Deadline.Format(TimeFormat),
				CreatedAt:   task.CreatedAt.Format(TimeFormat),
			},
		}, err
	}
}

func (s *KanbeeServiceServer) UpdateTaskState(ctx context.Context, in *pb.UpdateTaskStateRequest) (*pb.TaskResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		task, err := db.UpdateTaskState(uint(in.Id), in.State)
		if err != nil {
			return nil, err
		}
		return &pb.TaskResponse{
			Task: &pb.Task{
				Id:          int32(task.ID),
				ProjectId:   int32(task.ProjectID),
				Title:       task.Title,
				Description: task.Description,
				State:       task.State,
				Deadline:    task.Deadline.Format(TimeFormat),
				CreatedAt:   task.CreatedAt.Format(TimeFormat),
			},
		}, err
	}
}

func (s *KanbeeServiceServer) GetTaskById(ctx context.Context, in *pb.GetByIdRequest) (*pb.TaskResponse, error) {
	return nil, nil
}

func (s *KanbeeServiceServer) ListTasks(ctx context.Context, in *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		results, err := db.ListTask(uint(in.ProjectId), in.State)
		if err != nil {
			return nil, err
		}
		var tasks []*pb.Task
		for _, result := range results {
			task := &pb.Task{
				Id:          int32(result.ID),
				ProjectId:   int32(result.ProjectID),
				Title:       result.Title,
				Description: result.Description,
				State:       result.State,
				Deadline:    result.Deadline.Format(TimeFormat),
				CreatedAt:   result.CreatedAt.Format(TimeFormat),
			}
			tasks = append(tasks, task)
		}
		return &pb.ListTasksResponse{
			Tasks: tasks,
		}, err
	}
}

func (s *KanbeeServiceServer) ListUpdates(ctx context.Context, in *pb.ListUpdatesRequest) (*pb.ListUpdatesResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		results, err := db.ListUpdates(uint(in.TaskId))
		if err != nil {
			return nil, err
		}
		var updates []*pb.Update
		for _, result := range results {
			update := &pb.Update{
				Id:             int32(result.ID),
				TaskId:         int32(result.TaskID),
				Message:        result.Message,
				AttachmentData: result.AttachmentData,
				Filename:       result.Filename,
				CreatedAt:      result.CreatedAt.Format(TimeFormat),
			}
			updates = append(updates, update)
		}

		return &pb.ListUpdatesResponse{
			Updates: updates,
		}, nil
	}
}
