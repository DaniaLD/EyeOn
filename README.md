# Crypto Trading Gateway (Bitpin + nobitex)

یک پروژه Go برای اتصال به صرافی‌های ایرانی **بیت‌پین** و **نوبیتکس**، جهت ایجاد، لغو سفارش، دریافت موجودی کیف پول و مشاهده Order Book.

## Tech Stack

- Language: Go
- Web Framework: Gin
- Architecture: Hexagonal (Ports & Adapters)
- Exchanges supported: Bitpin , nobitex

---

## Hexagonal Architecture

این پروژه از **معماری هگزاگونال (Hexagonal Architecture)** پیروی می‌کند:

- `core/`: شامل business logic و interfaceها (ports)
- `adapters/api`: ورودی HTTP REST و هندلرهای مرتبط
- `adapters/exchange`: پیاده‌سازی آداپتر برای هر صرافی (Bitpin, nobitex)
- `cmd/main.go`: ورودی اصلی برنامه برای اجرای HTTP Server

این معماری باعث شده توسعه، تست و افزودن صرافی‌های جدید بسیار ساده و منعطف باشد.

