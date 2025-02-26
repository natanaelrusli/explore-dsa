class File {
  constructor(name, size) {
    this.name = name;
    this.size = size;
  }
}

class Directory {
  constructor(name) {
    this.name = name;
    this.contents = [];
  }

  add(item) {
    this.contents.push(item);
  }

  getSize() {
    let totalSize = 0;

    for (const item of this.contents) {
      if (item instanceof File) {
        totalSize += item.size;
      } else if (item instanceof Directory) {
        totalSize += item.getSize();
      }
    }

    return totalSize;
  }

  findFile(name) {
    for (const item of this.contents) {
      if (item instanceof File && item.name === name) {
        return item;
      } else if (item instanceof Directory) {
        const found = item.findFile(name);
        if (found) return found;
      }
    }

    return null;
  }
}

const root = new Directory("root");
const dir1 = new Directory("dir1");
const dir2 = new Directory("dir2");
const file1 = new File("file1.txt", 500);
const file2 = new File("file2.txt", 300);
const file3 = new File("file3.txt", 700);
const file4 = new File("file3.txt", 200);

dir1.add(file1);
dir1.add(file2);
dir2.add(file3);
root.add(dir1);
root.add(dir2);
dir2.add(file4);
console.log(dir1.findFile("file1.txt"));

console.log("Total Size:", root.getSize());
