class HotelElevator {
    public static void main(String... args) {
        int A[] = {40, 40, 100, 80, 20};
        int B[] = {3, 3, 2, 2, 3};
        int people = 5;
        int weight = 200;
        System.out.println(new HotelElevator().solution(A, B, people, weight));

        int X[] = {60, 80, 40};
        int Y[] = {2, 3, 5};
        people = 5;
        weight = 200;
        System.out.println(new HotelElevator().solution(X, Y, people, weight));
    }

    public int solution(int[] A, int[] B, int X, int Y) {
        Set<Integer> stops = new HashSet<>();

        int totalWeight = 0;
        int noOfPeople = 0;
        int noOfStop = 0;
        for (int i = 0; i < A.length; i++) {
            noOfPeople++;
            totalWeight += A[i];
            if (noOfPeople > X || totalWeight > Y) {
                noOfStop += stops.size();
                stops.clear();
                noOfPeople = totalWeight = 0;
                i--;
                noOfStop++;
            } else {
                stops.add(B[i]);
            }
        }
        return noOfStop + 1 + stops.size();
    }
}