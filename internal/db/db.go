package db

import (
	"time"

	"github.com/nrssi/kanbee/internal/db/models"
	pb "github.com/nrssi/kanbee/internal/services/kanban"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var KanbanStore *gorm.DB

func init() {
	var err error
	KanbanStore, err = gorm.Open(sqlite.Open("kanban.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	if err = KanbanStore.AutoMigrate(&models.Task{}, &models.Project{}, &models.Update{}); err != nil {
		panic("failed to migrate the schema to database")
	}
}

// Creates for all models
func CreateProject(name string, description string) (models.Project, error) {
	project := models.Project{
		Name:        name,
		Description: description,
		Tasks:       []models.Task{},
	}
	result := KanbanStore.Create(&project)
	return project, result.Error
}

func CreateTask(title string, description string, projectId uint, state pb.TaskState) (models.Task, error) {
	task := models.Task{
		Title:       title,
		Description: description,
		State:       state,
		Updates:     []models.Update{},
		ProjectID:   projectId,
	}
	result := KanbanStore.Create(&task)
	return task, result.Error
}

func CreateUpdate(taskId uint, message string, filename string, data []byte) (models.Update, error) {
	update := models.Update{
		CreatedAt:      time.Now(),
		Message:        message,
		Filename:       filename,
		AttachmentData: data,
		TaskID:         taskId,
	}
	result := KanbanStore.Create(&update)
	return update, result.Error
}

// Reads for all models
func ListProjects() ([]models.Project, error) {
	var projects []models.Project
	result := KanbanStore.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func ListTask(projectId uint, state pb.TaskState) ([]models.Task, error) {
	var tasks []models.Task
	query := "project_id = ?"
	var result *gorm.DB
	if len(state.String()) > 0 {
		query = query + " && state = ?"
		result = KanbanStore.Where(query, projectId, state).Find(&tasks)
	} else {
		result = KanbanStore.Where(query, projectId).Find(&tasks)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func ListUpdates(taskId uint) ([]models.Update, error) {
	var updates []models.Update
	result := KanbanStore.Where("task_id = ?", taskId).Find(&updates)
	if result.Error != nil {
		return nil, result.Error
	}
	return updates, nil
}

// Updates for all models
func UpdateProject(projectId uint, name string, description string) (models.Project, error) {
	var project models.Project
	result := KanbanStore.Where("project_id = ?", project).First(&project)
	if result.Error != nil {
		return models.Project{}, result.Error
	}
	project.Name = name
	project.Description = description
	KanbanStore.Save(&project)
	return project, nil
}

func UpdateTask(taskId uint, title string, description string, projectId uint, state pb.TaskState) (models.Task, error) {
	var task models.Task
	result := KanbanStore.Where("task_id = ?", taskId).First(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	task.Title = title
	task.Description = description
	task.ProjectID = projectId
	task.State = state
	KanbanStore.Save(&task)
	return task, nil
}

func UpdateTaskState(taskId uint, state pb.TaskState) (models.Task, error) {
	var task models.Task
	result := KanbanStore.Where("task_id = ?", taskId).First(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	task.State = state
	KanbanStore.Save(&task)
	return task, nil
}

func UpdateUpdate(updateId uint, taskId uint, message string, filename string, data []byte) (models.Update, error) {
	var update models.Update
	result := KanbanStore.Where("update_id = ? ", updateId).First(&update)
	if result.Error != nil {
		return models.Update{}, result.Error
	}
	update.Message = message
	update.AttachmentData = data
	update.Filename = filename
	update.TaskID = taskId
	KanbanStore.Save(&update)
	return update, nil
}

// Deletes for all models
func DeleteProject(projectId uint) (models.Project, error) {
	var project models.Project
	result := KanbanStore.Where("project_id = ? ", projectId).Delete(&project)
	if result.Error != nil {
		return models.Project{}, result.Error
	}
	return project, nil
}

func DeleteTask(taskId uint) (models.Task, error) {
	var task models.Task
	result := KanbanStore.Where("Task_id = ? ", taskId).Delete(&task)
	if result.Error != nil {
		return models.Task{}, result.Error
	}
	return task, nil
}

func DeleteUpdate(updateId uint) (models.Update, error) {
	var update models.Update
	result := KanbanStore.Where("Update_id = ? ", updateId).Delete(&update)
	if result.Error != nil {
		return models.Update{}, result.Error
	}
	return update, nil
}
