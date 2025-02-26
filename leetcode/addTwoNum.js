class ListNode {
  constructor(val = 0, next = null) {
    this.val = val;
    this.next = next;
  }
}
 
/**
 * @param {ListNode} l1 - The head of the first linked list.
 * @param {ListNode} l2 - The head of the second linked list.
 * @returns {ListNode} - The head of the linked list representing the sum.
 */
const addTwoNumbers = function(l1, l2) {
  const dummy = new ListNode();
  let carry = 0;
  let cur = dummy;
 
  while (l1 || l2 || carry) {
    const s = (l1?.val || 0) + (l2?.val || 0) + carry;
    carry = Math.floor(s / 10);
    cur.next = new ListNode(s % 10);
    cur = cur.next;
    l1 = l1?.next;
    l2 = l2?.next;
  }
 
  return dummy.next;
};

function printList(head) {
  const result = [];
  while (head) {
    result.push(head.val);
    head = head.next;
  }
  console.log(result.join(" -> "));
}

const l1 = new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9)))); // 9999
const l2 = new ListNode(1); // 1

printList(addTwoNumbers(l1, l2)); // 10000 -> ListNode(0 -> 0 -> 0 -> 0 -> 1))
