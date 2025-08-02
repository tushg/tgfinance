# TGF Finance Frontend

A modern React application for personal finance management built with TypeScript and Tailwind CSS.

## Features

- **Modern UI/UX**: Clean, responsive design with Tailwind CSS
- **TypeScript**: Full type safety and better developer experience
- **React Router**: Client-side routing for seamless navigation
- **React Query**: Efficient data fetching and caching
- **React Hook Form**: Performant form handling
- **Recharts**: Beautiful data visualization
- **Lucide React**: Modern icon library

## Tech Stack

- React 18
- TypeScript
- Tailwind CSS
- React Router DOM
- React Query
- React Hook Form
- Recharts
- Lucide React
- Axios

## Getting Started

### Prerequisites

- Node.js 18+
- npm or yarn

### Installation

```bash
# Install dependencies
npm install

# Start development server
npm start

# Build for production
npm run build

# Run tests
npm test
```

### Development

The application will be available at `http://localhost:3000`

### Environment Variables

Create a `.env` file in the frontend directory:

```env
REACT_APP_API_URL=http://localhost:8080
```

## Project Structure

```
frontend/
├── public/                 # Static files
├── src/
│   ├── components/         # Reusable components
│   │   ├── ui/            # Basic UI components
│   │   ├── layout/        # Layout components
│   │   ├── forms/         # Form components
│   │   └── charts/        # Chart components
│   ├── pages/             # Page components
│   ├── services/          # API services
│   ├── hooks/             # Custom React hooks
│   ├── types/             # TypeScript type definitions
│   ├── context/           # React context providers
│   ├── utils/             # Utility functions
│   └── assets/            # Static assets
├── package.json
├── tailwind.config.js
├── tsconfig.json
└── Dockerfile
```

## Available Scripts

- `npm start` - Start development server
- `npm build` - Build for production
- `npm test` - Run tests
- `npm eject` - Eject from Create React App

## Docker

```bash
# Build Docker image
docker build -t tgfinance-frontend .

# Run container
docker run -p 3000:3000 tgfinance-frontend
``` 