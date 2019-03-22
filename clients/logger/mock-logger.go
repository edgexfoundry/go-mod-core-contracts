/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package logger

// MockLogger is a type that can be used for mocking the LoggingClient interface during unit tests
type MockLogger struct {
}

// NewMockClient creates a mock instance of LoggingClient
func NewMockClient() LoggingClient {
	return MockLogger{}
}

// SetLogLevel simulates setting a log severity level
func (lc MockLogger) SetLogLevel(loglevel string) error {
	return nil
}

// Info simulates logging an entry at the INFO severity level
func (lc MockLogger) Info(msg string, args ...interface{}) {
}

// Debug simulates logging an entry at the DEBUG severity level
func (lc MockLogger) Debug(msg string, args ...interface{}) {
}

// Error simulates logging an entry at the ERROR severity level
func (lc MockLogger) Error(msg string, args ...interface{}) {
}

// Trace simulates logging an entry at the TRACE severity level
func (lc MockLogger) Trace(msg string, args ...interface{}) {
}

// Warn simulates logging an entry at the WARN severity level
func (lc MockLogger) Warn(msg string, args ...interface{}) {
}
