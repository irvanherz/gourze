# Gourze - Course Marketplace Backend

Gourze is a **high-performance course marketplace backend** built with **Golang**, designed to empower educators and learners by providing a scalable and efficient infrastructure. With a strong focus on modularity and performance, Gourze leverages **PostgreSQL** for reliable data storage, **GORM** as its ORM for seamless database interactions, and **Fx** for structured dependency injection. To ensure smooth media management, the platform integrates with **BunnyCDN**, enabling high-speed and cost-effective hosting of course videos and images. Gourze is built to handle real-world marketplace demands, making it an ideal solution for developers seeking experience in building modern, production-grade backend systems.

---

## 🚀 Features

- **Course Management**: Create, update, and manage online courses.
- **User Authentication**: Secure authentication system.
- **Media Hosting**: Integrates with **BunnyCDN** for storing videos and images.
- **Real-time Communication**: Uses **GORM** for database interactions.
- **Modular & Scalable Architecture**: Utilizes **Fx** for dependency injection.

---

## 🛠️ Tech Stack

- **Golang** - Backend logic
- **PostgreSQL** - Database
- **GORM** - ORM
- **Fx** - Dependency Injection
- **BunnyCDN** - Media hosting
- **Gin** - HTTP framework for REST API

---

## 📦 Installation

### **1. Clone the Repository**

```sh
git clone https://github.com/irvanherz/gourze.git
cd gourze
```

### **2. Set Up Environment Variables**

Create a `.env` file and add the following configurations:

```sh
DB_HOST=localhost
DB_USER=postgres
DB_PASS=xxx
DB_NAME=gourze
DB_PORT=5432

BUNNY_STORAGE_ZONE=xxx
BUNNY_STORAGE_ACCESS_KEY=xxx
BUNNY_STORAGE_REGION=sg
BUNNY_STORAGE_DOWNLOAD_BASE_URL=xxx

BUNNY_STREAM_ACCESS_KEY=xxx
BUNNY_STREAM_LIBRARY_ID=xxx
BUNNY_STREAM_UPLOAD_EXPIRATION_TIME=36000

JWT_SECRET=xxx
```

### **3. Install Dependencies**

```sh
go mod tidy
```

### **4. Setup Database Types**

Before running migrations, manually create required PostgreSQL enum types:

```sql
CREATE TYPE user_role AS ENUM ('super', 'admin', 'generic');
CREATE TYPE media_type AS ENUM ('image', 'document', 'video');
CREATE TYPE media_upload_status AS ENUM ('uploading','uploaded','processing','processed','failed');
CREATE TYPE order_status AS ENUM ('unpaid', 'paid', 'canceled');
```

### **5. Start the Server**

```sh
go run main.go
```

The backend should now be running on `http://localhost:8080`

---

## 📘 Documentation

For a complete list of API endpoints, request formats, and response structures, refer to the [API Documentation](https://github.com/yourusername/gourze/wiki).

---

## ✅ Contributing

1. Fork the repo
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m "Add new feature"`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a pull request

---

## 📝 License

This project is licensed under the **MIT License**.

---

## 🌟 Support

If you like this project, consider giving it a **star** ⭐ on GitHub!

