<?php

namespace App\Http\Controllers;

use App\Event;
use App\MyFavourites;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;


class MyFavouritesController extends Controller
{
    public function index(Request $request)
    {
        $myfavourites = MyFavourites::query();

        $userId = $request->query('user_id');

        $myfavourites->when($userId, function ($query) use ($userId) {
            return $query->where('user_id', '=', $userId);
        });
        return response()->json(['status' => 'success', 'data' => $myfavourites->paginate(20)]);
    }

    public function create(Request $request)
    {
        $rules = [
            'event_id' => 'required|integer',
            'user_id' => 'required|integer'
        ];

        $data = $request->all();

        $validator = Validator::make($data, $rules);

        if ($validator->fails()) {
            return response()->json([
                'status' => 'error',
                'message' => $validator->errors()
            ], 400);
        }

        $eventId = $request->input('event_id');
        $event  = Event::find($eventId);

        if (!$event) {
            return response()->json([
                'status' => 'error',
                'messagse' => 'event not found'
            ], 404);
        }

        $userId = $request->input('user_id');
        $user = getUser($userId);

        if ($user['status'] === 'error') {
            return response()->json([
                'status' => $user['status'],
                'message' => $user['message']
            ], $user['http_code']);
        }

        $isMyFavouritesExists = MyFavourites::where('event_id', '=', $eventId)
            ->where('user_id', '=', $userId)
            ->exists();

        if ($isMyFavouritesExists) {
            return response()->json([
                'status' => 'error',
                'message' => 'this event already set to favorite'
            ], 409);
        };

        $myfavourites = MyFavourites::create($data);
        return response()->json([
            'status' => 'success',
            'data' => $myfavourites
        ]);
    }

    public function destroy($id)
    {
        $myfavourites = MyFavourites::find($id);

        if (!$myfavourites) {
            return response()->json(['status' => 'error', 'message' => 'item not found'], 404);
        }

        $myfavourites->delete();
        return response()->json(['status' => 'success', 'message' => 'item deleted']);
    }
}
