var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

int requestCount = 0;

int requestsInCurrentSecond = 0;
double requestsPerSecond = 0;

// Timer to reset the counter every second
var timer = new System.Timers.Timer(1000); // 1 second interval
timer.Elapsed += (sender, e) => 
{
    // Set RPS to the number of requests in the last second
    requestsPerSecond = requestsInCurrentSecond;
    
    // Reset the counter for the next second
    requestsInCurrentSecond = 0;

    // Print the current RPS to the console
    Console.WriteLine($"Requests per Second (Current): {requestsPerSecond:F2}");
};
timer.Start();

app.MapGet("/", () => {
    requestsInCurrentSecond++;    

    return requestCount;
} );

app.Run();
