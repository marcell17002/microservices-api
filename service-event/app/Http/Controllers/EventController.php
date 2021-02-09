<?php

namespace App\Http\Controllers;

use App\Category;
use App\Event;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Validator;

class EventController extends Controller
{
    public function index(Request $request)
    {
        $event = Event::query();

        $q = $request->query('q');
        $title = $request->query('title');

        $event->when($q, function ($query) use ($q) {
            return $query->whereRaw("title LIKE '%" . strtolower($q) . "%'");
        });

        $event->when($title, function ($query) use ($title) {
            return $query->where('title', '=', $title);
        });
        return response()->json(['status' => 'success', 'data' => $event->paginate(20)]);
    }

    public function show($id)
    {
        $event = Event::find($id);
        if (!$event) {
            return response()->json(['status' => 'error', 'message' => 'event not found'], 404);
        }

        return response()->json(['status' => 'success', 'data' => $event]);
    }

    public function create(Request $request)
    {
        $rules = [
            'title' => 'required|string',
            'thumbnail' => 'string|url',
            'description' => 'required|string',
            'creator_id' => 'required|integer',
            'category_id' => 'required|integer',
        ];

        $data = $request->all();
        $validator = Validator::make($data, $rules);

        if ($validator->fails()) {
            return response()->json([
                'status' => 'error',
                'message' => $validator->errors()
            ], 400);
        }

        $categoryId = $request->input('category_id');
        $category = Category::find($categoryId);

        if (!$category) {
            return response()->json(['status' => 'error', 'message' => 'category not found']);
        }

        $event = Event::create($data);
        return response()->json(['status' => 'success', 'data' => $event]);
    }

    public function update(Request $request, $id)
    {
        $rules = [
            'title' => 'string',
            'thumbnail' => 'string|url',
            'description' => 'string',
            'creator_id' => 'integer',
            'category_id' => 'integer',
        ];

        $data = $request->all();
        $validator = Validator::make($data, $rules);

        if ($validator->fails()) {
            return response()->json([
                'status' => 'error',
                'message' => $validator->errors()
            ], 400);
        }

        $event = Event::find($id);
        if (!$event) {
            return response()->json([
                'status' => 'error',
                'message' => 'event not found'
            ], 404);
        }

        $categoryId = $request->input('category_id');
        if ($categoryId) {
            $category = Category::find($categoryId);
            if (!$category) {
                return response()->json([
                    'status' => 'error',
                    'message' => 'category not found'
                ], 404);
            }
        }

        $event->fill($data);
        $event->save();

        return response()->json(['status' => 'success', 'data' => $event]);
    }

    public function destroy($id)
    {
        $event = Event::find($id);

        if (!$event) {
            return response()->json(['status' => 'error', 'message' => 'event not found'], 404);
        }

        $event->delete();
        return response()->json(['status' => 'success', 'message' => 'event deleted']);
    }
}
