<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::get('category', 'CategoryController@index');
Route::get('category/{id}', 'CategoryController@show');
Route::post('category', 'CategoryController@create');
Route::put('category/{id}', 'CategoryController@update');
Route::delete('category/{id}', 'CategoryController@destroy');

Route::get('event', 'EventController@index');
Route::get('event/{id}', 'EventController@show');
Route::post('event', 'EventController@create');
Route::put('event/{id}', 'EventController@update');
Route::delete('event/{id}', 'EventController@destroy');

Route::get('myfavorites', 'MyFavouritesController@index');
Route::get('myfavorites/{id}', 'MyFavouritesController@show');
Route::post('myfavorites', 'MyFavouritesController@create');
Route::delete('myfavorites/{id}', 'MyFavouritesController@destroy');
//get by id
