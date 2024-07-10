<?php

$config = include('config.php');

// Настройки подключения к базе данных
$servername = $config['servername'];    // Сервер базы данных
$username = $config['username'];        // Имя пользователя базы данных
$password = $config['password'];        // Пароль пользователя базы данных
$dbname = $config['dbname'];            // Имя базы данных

try {
    // Подключение к базе данных
    $conn = mysqli_connect($servername, $username, $password, $dbname);

    // Проверка подключения
    if (!$conn) {
        die("Ошибка подключения: " . mysqli_connect_error());
    }

    // SQL-запрос для извлечения данных из таблицы Stat
    $sql = "SELECT id, root, size, elapced_time, date FROM stat";
    $result = $conn->query($sql);

    $sizes = [];
    $times = [];

    // Проверка, есть ли результаты
    if ($result->num_rows > 0) {
        echo "<table border='1'>";
        echo "<tr><th>ID</th><th>Root</th><th>Size</th><th>Time Spent</th><th>Date</th></tr>";

        while ($row = $result->fetch_assoc()) {
            $sizes[] = $row["size"];
            $times[] = $row["elapced_time"];
            echo "<tr>";
            echo "<td>" . $row["id"] . "</td>";
            echo "<td>" . $row["root"] . "</td>";
            echo "<td>" . $row["size"] . "</td>";
            echo "<td>" . $row["elapced_time"] . "</td>";
            echo "<td>" . $row["date"] . "</td>";
            echo "</tr>";
        }

        echo "</table>";
    } else {
        echo "0 results";
    }

    // Объединение размеров и времени в ассоциативный массив для сортировки
    $data = array_map(null, $sizes, $times);

    // Сортировка данных по размеру файла
    usort($data, function($a, $b) {
        return $a[0] - $b[0];
    });

    // Разделение данных обратно на два массива после сортировки
    $sizes = array_column($data, 0);
    $times = array_column($data, 1);

    // Закрытие подключения к базе данных
    $conn->close();
} catch(Exception $e) {
    echo "Ошибка: ". $e->getMessage();
}

?>


<!DOCTYPE html>
<html>
<head>
    <title>Chart.js Example</title>
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
</head>
<body>
    <div style="width: 100%; margin: auto;">
        <canvas id="myChart"></canvas>
    </div>

    <script>
        document.addEventListener('DOMContentLoaded', (event) => {
            const sizes = <?php echo json_encode($sizes); ?>;
            const times = <?php echo json_encode($times); ?>;

            const ctx = document.getElementById('myChart').getContext('2d');
            const myChart = new Chart(ctx, {
                type: 'line',
                data: {
                    labels: sizes,
                    datasets: [{
                        label: 'Зависимость размера файла от затраченного времени',
                        data: times,
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',
                        borderColor: 'rgba(75, 192, 192, 1)',
                        borderWidth: 1,
                        fill: false,
                        tension: 0.1
                    }]
                },
                options: {
                    scales: {
                        x: {
                            type: 'logarithmic',
                            position: 'bottom',
                            title: {
                                display: true,
                                text: 'Размер файла (байт)'
                            },
                            beginAtZero: true,
                        },
                        y: {
                            title: {
                                display: true,
                                text: 'Затраченное время (мс)'
                            },
                            beginAtZero: true,
                            ticks: {
                                stepSize: 1
                            }
                        }
                    }
                }
            });
        });
    </script>
</body>
</html>