# OpenAPI

## Định nghĩa
OpenAPI là một đặc tả để xác định và ghi tài liệu cho REST API. OpenAPI cho phép các nhà phát triển mô tả cấu trúc API ở dạng chuẩn hóa hơn, làm cho việc thiết kế, xây dựng và sử dụng API trở nên dễ dàng hơn.

## Đặc điểm
* OpenAPI cung cấp 1 ngôn ngữ chung cho API, đảm bảo tính đồng nhất của API trên các nền tảng khác nhau.
* Đặc tả thường được viết bằng JSON hoặc YAML, cho phép khả năng tự động hóa và tích hợp với đa dạng các công cụ khác.
* Các công cụ như Swagger UI có thể mô phỏng đặc tả OpenAPI thành tài liệu API cho phép người dùng tương tác trực tiếp với API đó.
* OpenAPI cung cấp khả năng sinh code từ đặc tả API.
* Hỗ trợ xác thực và kiểm tra tự động các yêu cầu và phản hồi của API.

## Ưu điểm
* Tăng khả năng kết hợp giữa devs, testers và product managers.
* Nâng cao khả năng khám phá và sử dụng API.
* Giảm thời gian phát triển thông qua tự động hóa và tạo mã.
* Hỗ trợ quản trị API và kiểm soát phiên bản tốt hơn.

##  Cấu trúc cơ bản của OpenAPI
* **info**: API metadata, bao gồm title, version và description.
* **servers**: URL của server và môi trường như local, staging hay product.
* **paths**: Định nghĩa endpoint và phương thức HTTP được hỗ trợ (ví dụ: GET, POST, PUT, DELETE).
* **components**: Các thành phần có thể tái sử dụng như schema, parameter và response.

## Ví dụ:
```yaml
openapi: 3.1.0
info:
  title: Swagger Example APIs
  description: Simple APIs written in Go to demonstrate Swagger
  version: 1.0.0
servers:
  - url: http://localhost:1906
    description: Local server
paths:
  /users:
    get:
      summary: Get all users
      description: Returns a list of all users
      responses:
        '200':
          description: A list of users
          content:
            application/json:
              schema:
                type: array
                items:
                  type: object
                  properties:
                    id:
                      type: integer
                    username:
                      type: string
                    name:
                      type: string
        '500':
			description: Internal server error
```
File openapi.yaml trên định nghĩa server gồm http://localhost:1906 là server nội bộ. Các endpoint của server bao gồm:

* '/users': Trả về danh sách tất cả các users. Status code:
	* 200: Chạy thành công và trả về 1 JSON array, với mỗi phần tử là 1 JSON gồm các thuộc tính và kiểu dữ liệu: id integer, username string và name string.
	* 500: Server lỗi.