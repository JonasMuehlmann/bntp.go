// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var ConfigValidator = validator.New()

const (
	ValidatorLogrusLogLevel            = "logrus_log_level"
	ValidateDBDriver                   = "db_driver"
	ValidatorBookmarkManager           = "bookmark_manager"
	ValidatorTagsManager               = "tags_manager"
	ValidatorDocumentManager           = "document_manager"
	ValidatorDocumentContentManager    = "document_content_manager"
	ValidatorBookmarkRepository        = "bookmark_repository"
	ValidatorTagsRepository            = "tags_repository"
	ValidatorDocumentRepository        = "document_repository"
	ValidatorDocumentContentRepository = "document_content_repository"
	ValidatorHooks                     = "hooks"
	ValidatorHookPoint                 = "hook_point"
	ValidatorHook                      = "hook"
)

// TODO: Add test to validate that The field name, tag "name" and "mapstructure" match.
type Config struct {
	LogFile         string        `name:"log_file" mapstructure:"log_file" validate:"required,file"`
	ConsoleLogLevel string        `name:"console_log_level" mapstructure:"console_log_level" validate:"required,logrus_log_level"`
	FileLogLevel    string        `name:"file_log_level" mapstructure:"file_log_level" validate:"required,logrus_log_level"`
	DB              DBConfig      `name:"db" mapstructure:"db" validate:"required"`
	Backend         BackendConfig `name:"backend" mapstructure:"backend" validate:"required"`
}

type DBConfig struct {
	Driver     string   `name:"driver" mapstructure:"driver" validate:"required,db_driver"`
	DataSource string   `name:"data_source" mapstructure:"data_source" validate:"required,file"`
	Args       []string `name:"args" mapstructure:"args"`
}

type BackendConfig struct {
	Bookmarkmanager        BookmarkManagerConfig        `name:"bookmark_manager" mapstructure:"bookmark_manager" validate:"required,bookmark_manager"`
	TagsManager            TagsManagerConfig            `name:"tags_manager" mapstructure:"tags_manager" validate:"required,tags_manager"`
	DocumentManager        DocumentManagerConfig        `name:"document_manager" mapstructure:"document_manager" validate:"required,document_manager"`
	DocumentContentManager DocumentContentManagerConfig `name:"document_content_manager" mapstructure:"document_content_manager" validate:"required,document_content_manager"`
}

// ******************************************************************//
//                        Repository configs                         //
// ******************************************************************//

type BookmarkRepositoryConfig struct {
	DB            DBConfig             `name:"db" mapstructure:"db" validate:"required"`
	TagRepository TagsRepositoryConfig `name:"tag_repository" mapstructure:"tag_repository" validate:"required"`
}

type TagsRepositoryConfig struct {
	DB DBConfig `name:"db" mapstructure:"db" validate:"required"`
}

type DocumentRepositoryConfig struct {
	DB            DBConfig             `name:"db" mapstructure:"db" validate:"required"`
	TagRepository TagsRepositoryConfig `name:"tag_repository" mapstructure:"tag_repository" validate:"required"`
}

type DocumentContentRepositoryConfig struct {
	DB DBConfig `name:"db" mapstructure:"db" validate:"required"`
}

type HooksConfig struct {
	HookPoints map[string][]string `name:"hooks" mapstructure:"hooks" validate:"dive,keys,hook_point,endkeys,dive,hook"`
}

// ******************************************************************//
//                          Manager configs                          //
// ******************************************************************//

type BookmarkManagerConfig struct {
	Hooks              HooksConfig              `name:"hooks" mapstructure:"hooks"`
	BookmarkRepository BookmarkRepositoryConfig `name:"bookmark_repository" mapstructure:"bookmark_repository" validate:"required,bookmark_repository"`
}

type TagsManagerConfig struct {
	Hooks          HooksConfig          `name:"hooks" mapstructure:"hooks"`
	TagsRepository TagsRepositoryConfig `name:"tags_repository" mapstructure:"tags_repository" validate:"required,tags_repository"`
}

type DocumentManagerConfig struct {
	Hooks              HooksConfig              `name:"hooks" mapstructure:"hooks"`
	DocumentRepository DocumentRepositoryConfig `name:"document_repository" mapstructure:"document_repository" validate:"required,document_repository"`
}

type DocumentContentManagerConfig struct {
	Hooks                     HooksConfig                     `name:"hooks" mapstructure:"hooks"`
	DocumentContentRepository DocumentContentRepositoryConfig `name:"document_content_repository" mapstructure:"document_content_repository" validate:"required,document_content_repository"`
}

func init() {
	err := ConfigValidator.RegisterValidation(ValidatorLogrusLogLevel, validateLogLevel)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidateDBDriver, validateDBDriver)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidatorBookmarkManager, validate_bookmark_manager)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidatorTagsManager, validate_tags_manager)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidatorDocumentManager, validate_document_manager)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidatorDocumentContentManager, validate_document_content_manager)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidatorBookmarkRepository, validate_bookmark_repository)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidatorTagsRepository, validate_tags_repository)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidatorDocumentRepository, validate_document_repository)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidatorDocumentContentRepository, validate_document_content_repository)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidatorHookPoint, validate_hook_point)
	if err != nil {
		panic(err)
	}

	err = ConfigValidator.RegisterValidation(ValidatorHook, validate_hook)
	if err != nil {
		panic(err)
	}
}

// ******************************************************************//
//                          Custom validators                        //
// ******************************************************************//

func validateLogLevel(field validator.FieldLevel) bool {
	_, err := logrus.ParseLevel(field.Field().String())

	return err == nil
}

func validateDBDriver(field validator.FieldLevel) bool {
	return field.Field().String() == "sqlite3"
}

func validate_hook(field validator.FieldLevel) bool {
	// check if name in predefined or registered hooks

	return false
}

func validate_hook_point(field validator.FieldLevel) bool {
	// Check if name in hook points enum

	return false
}

// ********************    Manager validators    ********************//

func validate_bookmark_manager(field validator.FieldLevel) bool {
	bookmarkManagerConfig, ok := field.Field().Interface().(BookmarkManagerConfig)
	if !ok {
		return false
	}

	err := ConfigValidator.Struct(bookmarkManagerConfig)
	if err != nil {
		return false
	}

	hooks, ok := field.Field().Interface().(HooksConfig)
	if !ok {
		return false
	}

	err = ConfigValidator.Struct(hooks)

	return err == nil
}

func validate_tags_manager(field validator.FieldLevel) bool {
	tagsManagerConfig, ok := field.Field().Interface().(TagsManagerConfig)
	if !ok {
		return false
	}

	err := ConfigValidator.Struct(tagsManagerConfig)
	if err != nil {
		return false
	}

	hooks, ok := field.Field().Interface().(HooksConfig)
	if !ok {
		return false
	}

	err = ConfigValidator.Struct(hooks)

	return err == nil
}

func validate_document_manager(field validator.FieldLevel) bool {
	DocumentManagerConfig, ok := field.Field().Interface().(DocumentManagerConfig)
	if !ok {
		return false
	}

	err := ConfigValidator.Struct(DocumentManagerConfig)
	if err != nil {
		return false
	}

	hooks, ok := field.Field().Interface().(HooksConfig)
	if !ok {
		return false
	}

	err = ConfigValidator.Struct(hooks)

	return err == nil
}

func validate_document_content_manager(field validator.FieldLevel) bool {
	documentContentManagerConfig, ok := field.Field().Interface().(DocumentContentManagerConfig)
	if !ok {
		return false
	}

	err := ConfigValidator.Struct(documentContentManagerConfig)
	if err != nil {
		return false
	}

	hooks, ok := field.Field().Interface().(HooksConfig)
	if !ok {
		return false
	}

	err = ConfigValidator.Struct(hooks)

	return err == nil
}

// *******************    Repository validators    ******************//

func validate_bookmark_repository(field validator.FieldLevel) bool {
	bookmarkRepositoryConfig, ok := field.Field().Interface().(BookmarkRepositoryConfig)
	if !ok {
		return false
	}

	err := ConfigValidator.Struct(bookmarkRepositoryConfig)

	return err == nil
}

func validate_tags_repository(field validator.FieldLevel) bool {
	tagsRepositoryConfig, ok := field.Field().Interface().(TagsRepositoryConfig)
	if !ok {
		return false
	}

	err := ConfigValidator.Struct(tagsRepositoryConfig)

	return err == nil
}

func validate_document_repository(field validator.FieldLevel) bool {
	DocumentRepositoryConfig, ok := field.Field().Interface().(DocumentRepositoryConfig)
	if !ok {
		return false
	}

	err := ConfigValidator.Struct(DocumentRepositoryConfig)

	return err == nil
}

func validate_document_content_repository(field validator.FieldLevel) bool {
	documentContentRepositoryConfig, ok := field.Field().Interface().(DocumentContentRepositoryConfig)
	if !ok {
		return false
	}

	err := ConfigValidator.Struct(documentContentRepositoryConfig)

	return err == nil
}
