class ParkingSystem {

    public int[] parkingList;

    public ParkingSystem(int big, int medium, int small) {
        this.parkingList = new int[]{big, medium, small};
    }
    
    public boolean addCar(int carType) {
        if (this.parkingList[carType - 1] > 0) {
            this.parkingList[carType - 1] -= 1;
            return true;
        }
        return false;
    }
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * ParkingSystem obj = new ParkingSystem(big, medium, small);
 * boolean param_1 = obj.addCar(carType);
 */
