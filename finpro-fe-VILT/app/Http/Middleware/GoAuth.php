<?php

namespace App\Http\Middleware;

use App\Services\GoApiService;
use Closure;
use Illuminate\Http\Request;
use Symfony\Component\HttpFoundation\Response;

class GoAuth
{
    public function __construct(protected GoApiService $api) {}

    public function handle(Request $request, Closure $next): Response
    {
        if (!$this->api->isLoggedIn()) {
            return redirect()->route('login');
        }

        return $next($request);
    }
}
