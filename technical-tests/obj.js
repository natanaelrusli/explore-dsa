var a = 20;

var obj = {
  a: 10,
  say: (a) => {
    console.log(a);
  },
  say2: function () {
    console.log(this.a);
  },
};

obj.say(obj.a);
obj.say2(obj.a);
