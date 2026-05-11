<?php

namespace App\Http\Controllers;

use App\Services\GoApiService;
use Inertia\Inertia;
use Inertia\Response;

class DashboardController extends Controller
{
    public function __construct(protected GoApiService $api) {}

    public function index(): Response
    {
        $overview = $this->api->getDashboard();
        $data = $overview['data'] ?? [];

        // Last assessment result — prefer session cache if fresh, fallback to API
        $result = session('onboarding_result') ?: ($data['last_assessment'] ?? []);

        return Inertia::render('Dashboard', [
            'dashboard' => $data,
            'result'    => $result,
        ]);
    }
}
