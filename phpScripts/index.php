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

    // Получение данных из POST-запроса
    $data = json_decode(file_get_contents('php://input'), true);

    // Проверка наличия необходимых полей в массиве данных
    if (isset($data['root']) && isset($data['size']) && isset($data['timeSpent'])) {
        $root = $data['root'];
        $size = $data['size'];
        $time_spent = $data['timeSpent'];

        $stmt = $conn->prepare("INSERT INTO stat (root, size, elapced_time) VALUES (?, ?, ?)");
        $stmt->bind_param("sii", $root, $size, $time_spent);

        if ($stmt->execute()) {
            echo json_encode(["status" => "success"]);
        } else {
            echo json_encode(["status" => "error", "message" => $stmt->error]);
        }

        $stmt->close();
    } else {
        echo json_encode(["status" => "error", "message" => "Invalid input"]);
    }

    // закрытие соединения с базой данных
    mysqli_close($conn);
} catch(Exception $e) {
    echo "Ошибка: ". $e->getMessage();
}

?>
