package prototype

func ExampleFile() {
	file1 := &file{name: "file1"}
	cloneFile := file1.clone()
	cloneFile.print("example_")
	// Output:
	// example_file1_clone
}

func ExampleFolder() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}
	folder1 := &folder{
		name:     "folder1",
		children: []inode{file1, file2, file3},
	}
	cloneFolder := folder1.clone()
	cloneFolder.print("example_")
	// Output:
	// example_folder1_clone
	// example_example_file1_clone
	// example_example_file2_clone
	// example_example_file3_clone
}
