import Link from 'next/link';

export default function NotFound() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50">
      <div className="text-center">
        <h1 className="text-6xl font-bold text-gray-900 mb-4">404</h1>
        <h2 className="text-2xl font-semibold text-gray-700 mb-6">Page Not Found</h2>
        <p className="text-gray-600 mb-8">The page you are looking for doesn&apos;t exist or has been moved.</p>
        <Link
          href="/"
          className="px-6 py-3 bg-[#ff8c42] text-white rounded-lg hover:bg-[#ff8c42]/90 transition-colors"
        >
          Go back home
        </Link>
      </div>
    </div>
  );
} 