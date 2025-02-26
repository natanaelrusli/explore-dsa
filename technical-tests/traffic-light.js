const colors = ["red", "yellow", "green"];

async function printColors() {
  for await (const color of colors) {
    console.log(color);
  }
}

// printColors();

// Alternative delay function using a synchronous loop and microtasks
async function delayAlternative(ms) {
  const start = performance.now();
  while (performance.now() - start < ms) {
    // Do nothing, let the microtask queue process
    await Promise.resolve();
  }
}

delayAlternative(200);
