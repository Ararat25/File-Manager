<?php

// Настройки подключения к базе данных
$servername = "localhost";  // Сервер базы данных
$username = "root";         // Имя пользователя базы данных
$password = "Vaa65770407!"; // Пароль пользователя базы данных
$dbname = "RBS";            // Имя базы данных

// Подключение к базе данных
$conn = mysqli_connect($servername, $username, $password, $dbname);

// Проверка подключения
if (!$conn) {
    die("Ошибка подключения: " . mysqli_connect_error());
}

// SQL-запрос для извлечения данных из таблицы Stat
$sql = "SELECT * FROM stat";
$result = $conn->query($sql);

// Проверка, есть ли результаты
if ($result->num_rows > 0) {
    echo "<table border='1'>";
    echo "<tr><th>ID</th><th>Root</th><th>Size</th><th>Time Spent</th></tr>";

    while ($row = $result->fetch_assoc()) {
        echo "<tr>";
        echo "<td>" . $row["id"] . "</td>";
        echo "<td>" . $row["root"] . "</td>";
        echo "<td>" . $row["size"] . "</td>";
        echo "<td>" . $row["elapced_time"] . "</td>";
        echo "</tr>";
    }

    echo "</table>";
} else {
    echo "0 results";
}

// Закрытие подключения к базе данных
$conn->close();
?>
