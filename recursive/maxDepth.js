class TreeNode {
  constructor(val) {
    this.val = val;
    this.left = null;
    this.right = null;
  }
}

function maxDepth(root) {
  if (root === null) {
    return 0;
  }

  return 1 + Math.max(maxDepth(root.left), maxDepth(root.right));
}

// Example tree:
//        1
//       / \
//      2   3
//     / \    
//    4   5   
const root = new TreeNode(1);
root.left = new TreeNode(2);
root.right = new TreeNode(3);
root.left.left = new TreeNode(4);
root.left.right = new TreeNode(5);
root.left.right.right = new TreeNode(5);

console.log(maxDepth(root));
