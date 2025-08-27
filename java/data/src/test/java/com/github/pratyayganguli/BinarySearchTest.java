package com.github.pratyayganguli;

import com.github.pratyayganguli.search.BinarySearch;
import org.junit.jupiter.api.Test;

public class BinarySearchTest {

    @Test
    void search() {
        BinarySearch bs = new BinarySearch();
        int[] input = {-50, 1, 2000, 3223, 91232, 6123002};
        bs.search(input, 3223);
    }
}
