// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file generates the volume of a pyramid 

"use strict"

function calculate() {
  // input
  const width = parseInt(document.getElementById("width").value)
  const height = parseInt(document.getElementById("height").value)
  const length = parseInt(document.getElementById("length").value)

  // process
  const volume = length * width * height / 3 

  // output
  document.getElementById("volume").innerHTML = "Volume is: " + volume + " cm³"
}
