package com.github.pratyayganguli.search;

public class BinarySearch {

    public void search(int [] input, int key) {
        int res = recursiveCall(0, input.length-1, key, input);
        if (res == -1)  {
            System.out.println("element does not exist!");
        } else {
            System.out.println(res);
        }
    }

    private int recursiveCall(int low, int high, int key, int[] arr) {
        while(low<=high) {
            int mid = low + (high-low) / 2;
            if (arr[mid] < key) {
                low = mid + 1;
                return recursiveCall(low, high, key, arr);
            } else if (arr[mid] > key) {
                high = mid - 1;
                return recursiveCall(low, high, key, arr);
            } else {
                return mid;
            }
        }
        // element not found
        return -1;
    }
}
