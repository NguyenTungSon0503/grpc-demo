const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');

const PROTO_PATH = path.join(__dirname, '../../proto/helloworld.proto');
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

function createGrpcClient(serviceName) {
  const { helloworld } = grpc.loadPackageDefinition(packageDefinition);
  return new helloworld[serviceName](
    'localhost:50051',
    grpc.credentials.createInsecure()
  );
}

function sayHello(client, request) {
  client.sayHello(request, (err, response) => {
    if (err) {
      console.error('Error:', err);
      return;
    }
    console.log('Response from server (Message):', response.message);
    console.log('Response from server (Age):', response.age);
  });
}

function calculateSummary(client, request) {
  client.Summary(request, (err, response) => {
    if (err) {
      console.error('Error:', err);
      return;
    }
    console.log('Response from server (Sum):', response.sum);
  });
}

const clientCalculate = createGrpcClient('Calculate');
const clientHello = createGrpcClient('Greeter');

const requestSummary = { num1: 10, num2: 12 };
const requestHello = { name: 'Son', age: '20' };

sayHello(clientHello, requestHello);
calculateSummary(clientCalculate, requestSummary);
