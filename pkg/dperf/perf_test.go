// This file is part of MinIO dperf
// Copyright (c) 2021-2025 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package dperf

import (
	"context"
	"os"
	"runtime"
	"testing"
)

func TestDrivePerfRun(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "dperf-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	perf := &DrivePerf{
		Serial:     true,
		BlockSize:  4096,
		FileSize:   4096 * 4,
		IOPerDrive: 1,
		SyncMode:   true,
	}

	results, err := perf.Run(context.Background(), tmpDir)
	if runtime.GOOS != "linux" {
		if err == nil && (len(results) == 0 || results[0].Error == nil) {
			t.Errorf("expected error on non-linux platform, got nil")
		}
		return
	}

	if err != nil {
		t.Fatalf("unexpected error running dperf: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Error != nil {
		t.Fatalf("result error: %v", results[0].Error)
	}
}
