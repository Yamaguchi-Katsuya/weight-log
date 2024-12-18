
# Weight Log Visualization System

This system is a web application that visualizes weight logs using interactive line charts. The application is built with Go and leverages the `go-echarts` library for chart rendering.

## Features

- Visualizes weight logs with an interactive line chart.
- Smooth line representation of the weight data.
- Tooltip support for data points.
- Fully responsive chart layout.
- Supports a customizable Y-axis range and grid styling.

## Requirements

- Go (1.18 or later)
- A JSON file containing weight log data (e.g., `weight-log.json`).

## Setup Instructions

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. Copy the example JSON file and rename it:

   ```bash
   cp weight-log.example.json weight-log.json
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

5. Open your browser and navigate to:

   ```
   http://localhost:8082
   ```

## File Structure

- `main.go`: The main application entry point.
- `weight-log.example.json`: Example JSON file for weight log data.
- `weight-log.json`: Actual JSON file used by the application (not included in version control).

## JSON File Format

The `weight-log.json` file must have the following structure:

```json
{
  "weight_logs": [
    {
      "date": "24-01-01",
      "weight": 65.0
    },
    {
      "date": "24-01-02",
      "weight": 64.8
    }
  ]
}
```

- `date`: Date of the weight log in `YY-MM-DD` format.
- `weight`: Weight in kilograms (floating-point number).

## Dependencies

- `github.com/go-echarts/go-echarts/v2`: Library for creating interactive charts.

## Customization

You can customize the chart settings by modifying the following sections in `main.go`:

- **Y-Axis Range:** Adjust `Min` and `Max` in `WithYAxisOpts`.
- **Grid Layout:** Adjust `Top`, `Left`, `Right`, and `Bottom` in `WithGridOpts`.
- **Tooltip:** Modify the tooltip formatter in `WithTooltipOpts`.

## License

This project is licensed under the MIT License.
