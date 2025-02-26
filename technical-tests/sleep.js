function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

function trafficLight() {
  sleep(0).then(() => console.log("Red"));
  sleep(1000).then(() => console.log("Yellow"));
  sleep(2000).then(() => console.log("Green"));
  sleep(4000).then(() => trafficLight());
}

async function trafficLightAsync() {
  await sleep(0);
  console.log("red");
  await sleep(1000);
  console.log("yellow");
  await sleep(2000);
  console.log("green");
  await sleep(4000);
  trafficLightAsync();
}

trafficLightAsync();
