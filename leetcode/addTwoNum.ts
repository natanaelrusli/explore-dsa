interface ItemNodeI {
  val: number;
  next: ItemNode | null;
}

class ItemNode implements ItemNodeI {
  val: number;
  next: ItemNode | null;

  constructor(val: number = 0, next: ItemNode | null = null) {
    this.val = val;
    this.next = next;
  }
}

function addTwoNumbersTs(l1: ItemNode | null, l2: ItemNode | null): ItemNode | null {
  const dummy: ItemNodeI = new ItemNode();
  let carry: number = 0;
  let cur: ItemNodeI = dummy;

  while (l1 || l2 || carry) {
    const s: number = (l1?.val || 0) + (l2?.val || 0) + carry;
    carry = Math.floor(s / 10);
    
    cur.next = new ItemNode(s % 10);
    cur = cur.next;
    l1 = l1?.next || null;
    l2 = l2?.next || null;
  }

  return dummy.next;
}

function printListTs(head: ItemNode | null) {
  const result: number[] = [];
  while (head) {
    result.push(head.val);
    head = head.next;
  }

  console.log(result.join(" --> "));
}

const list1 = new ItemNode(1);
const list2 = new ItemNode(9, new ItemNode(9, new ItemNode(9, new ItemNode(9))));

printListTs(addTwoNumbersTs(list1, list2));
